#!/bin/bash

if [[ $# -ne 1 ]]; then
	echo "Usage:"
	echo "$0 {cache server number}"
	exit 1
fi

cs_num=$1

# TODO: should also set upperbound next year.
[[ $cs_num -le 2 ]] && {
	echo "Error: cache server should be more than 3 ($cs_num provided)"
	exit 2
}

which jq >/dev/null 2>&1 || {
	echo "Error: please install jq first."
	exit 3
}

PORT_BASE=9526
HOST_BASE=127.0.0.1
MAX_ITER=500
DELETED_KEYS=()
_DELETED_KEYS_GENERATED=0

PASS_PROMPT="\e[1;32mPASS\e[0m"
FAIL_PROMPT="\e[1;31mFAIL\e[0m"

#生成一个随机端口的 URL，并且该端口是在 $PORT_BASE 的基础上加上一个在 1 到 $cs_num 之间的随机数
function get_cs() {
	port=$(($PORT_BASE + $(shuf -i 1-$cs_num -n 1)))
	echo http://$HOST_BASE:$port
}

#随机生成一个"key-随机数"值，在1到MAX_ITER=500之间
function get_key() {
	echo "key-$(shuf -i 1-$MAX_ITER -n 1)"
}

#生成一组键值，形式为"key-随机数",随机数在1-count之间，并将其添加到DELETED_KEYS数组中，=~是正则表达式匹配操作符
function gen_deleted_keys() {
	[[ $_DELETED_KEYS_GENERATED == 1 ]] && return 0

	local count=$((MAX_ITER / 10 * 3))

	while [[ ${#DELETED_KEYS[@]} -lt $count ]]; do
		local key="key-$(shuf -i 1-$MAX_ITER -n 1)"
		if ! [[ " ${DELETED_KEYS[@]} " =~ " ${key} " ]]; then
			DELETED_KEYS+=("$key")
		fi
	done

	_DELETED_KEYS_GENERATED=1
}

#生成一个包含索引信息的 JSON 对象，其中键是 "key-索引值"，值是 "value 索引值"。这个 JSON 对象通过 jq 工具生成，并最终输出到标准输出。
function gen_json_with_idx() {
	local idx=$1

	jq -n --arg key "key-$idx" --arg value "value $idx" '{($key): ($value)}'
}

#通过解析传入的键 $key 来获取索引值，然后调用另一个函数 gen_json_with_idx 来生成包含该索引值的 JSON 对象。
function gen_json_with_key() {
	local idx=$(echo $key | cut -d- -f2)

	gen_json_with_idx $idx
}

#比较两个 JSON 对象中指定键的值是否相等，它通过使用 jq 解析 JSON 数据，提取指定键的值，然后进行比较。
function compare_json_for_key() {
	local key=$1
	local result=$2
	local expect=$3

	local value1=$(echo "$result" | jq -r ".\"$key\"" 2>/dev/null)
	local value2=$(echo "$expect" | jq -r ".\"$key\"" 2>/dev/null)

	[[ "$value1" = "$value2" ]]
}

#用于查询指定键的值，然后根据期望的存在性和期望的结果，验证 HTTP 状态码和查询结果是否符合预期。如果符合预期，函数返回值为0（成功）；否则，返回值为非零（失败）。
function query_key() {
	local key=$1
	local exist=$2
	#这一行使用 curl 命令发起 HTTP GET 请求，查询指定键的值。response 变量包含了从指定 URL 收到的 HTTP 响应的内容以及最后一行的 HTTP 状态码。
	local response=$(curl -s -w "\n%{http_code}" $(get_cs)/$key)
	# everything but the last line. `head -n -1` breaks in macos, let's turn to sed trick.
	#删除最后一行返回的状态码
	local result=$(echo "$response" | sed '$d')
	#提取最后一行返回的状态码
	local status_code=$(echo "$response" | tail -n 1)
	#基于是否存在分别检测，如果 exist 的值为 1，表示期望键存在
	if [[ $exist == 1 ]]; then
		local expect=$(gen_json_with_key $key)
		if [[ $status_code -ne 200 ]] || ! compare_json_for_key "$key" "$result" "$expect"; then
			echo -e "Error:\tInvalid response"
			echo -e "\texpect: 200 $expect"
			echo -e "\tgot: $status_code $result"
			return 1
		fi
	else
		if [[ $status_code -ne 404 ]]; then
			echo "Error: expect status code 404 but got $status_code"
			return 1
		fi
	fi
}

#执行一系列 HTTP POST 请求，向指定的 URL 发送带有特定索引的 JSON 数据，然后验证每次请求的 HTTP 状态码是否为 200。如果任何一次请求的状态码不是 200，函数就会输出错误信息并返回非零退出码。
function test_set() {
	local i=1

	while [[ $i -le $MAX_ITER ]]; do
		status_code=$(curl -s -o /dev/null -w "%{http_code}" -XPOST -H "Content-type: application/json" -d "$(gen_json_with_idx $i)" $(get_cs))
		if [[ $status_code -ne 200 ]]; then
			echo "Error: expect status code 200 but got $status_code"
			return 1
		fi
		((i++))
	done
}

#这个函数的作用是执行一系列 HTTP GET 请求，向指定的 URL 发送请求，并验证每次请求的 HTTP 状态码是否为 200。期望查询的键是存在的，用于查询检查。
function test_get() {
	local count=$((MAX_ITER / 10 * 3))
	local i=0

	while [[ $i -lt $count ]]; do
		query_key $(get_key) 1 || return 1
		((i++))
	done
}

#这个函数的作用是执行一系列 HTTP DELETE 请求，删除指定键的数据，并验证每次请求的 HTTP 状态码是否为 200，且响应内容是否符合预期。根据要求delete操作删除已存在键值对时返回状态码200，返回内容为1。
function test_delete() {
	gen_deleted_keys
	for key in "${DELETED_KEYS[@]}"; do
		local response=$(curl -XDELETE -s -w "\n%{http_code}" $(get_cs)/$key)
		# `head -n 1` works for delete actually. let's use sed for consistency.
		local result=$(echo "$response" | sed '$d')
		local status_code=$(echo "$response" | tail -n 1)
		local expect=1
		if [[ $status_code -ne 200 ]] || [[ "$result" != "$expect" ]]; then
			echo -e "Error:\tInvalid response"
			echo -e "\texpect: $status_code $expect"
			echo -e "\tgot: $status_code $result"
			return 1
		fi
	done
}

# need to check all keys to guarantee only appointed keys are removed.
#这个函数的作用是在执行了一系列的删除操作后，验证一系列键是否仍然存在。
function test_get_after_delete() {
	local key
	local exist
	local i=1
	while [[ $i -le $MAX_ITER ]]; do
		key=$(get_key)
		#若在删除数组中则给exist赋值为0否则赋值为1
		[[ " ${DELETED_KEYS[@]} " =~ " ${key} " ]] && exist=0 || exist=1

		query_key $key $exist || return 1

		((i++))
	done
}

#这个函数的作用是再次执行一系列 HTTP DELETE 请求，删除指定键的数据，并验证每次请求的 HTTP 状态码是否为 200，且响应内容是否符合预期，即为0。
function test_delete_after_delete() {
	for key in "${DELETED_KEYS[@]}"; do
		local response=$(curl -XDELETE -s -w "\n%{http_code}" $(get_cs)/$key)
		local result=$(echo "$response" | sed '$d')
		local status_code=$(echo "$response" | tail -n 1)
		#若要删除的键值对已经被删除则返回状态码200及内容0
		if [[ $status_code -ne 200 ]] || [[ "$result" != "0" ]]; then
			echo -e "Error:\tInvalid response"
			echo -e "\texpect: 200 0"
			echo -e "\tgot: $status_code $result"
			return 1
		fi
	done
}

#这个函数的作用是运行传递给它的测试函数，并根据测试函数的执行结果输出相应的消息。如果测试函数成功执行，输出成功的提示；如果测试函数执行失败，输出失败的提示。
function run_test() {
	local test_function=$1
	local test_name=$2

	# echo "starting $test_name..."
	if $test_function; then
		echo -e "$test_name ...... ${PASS_PROMPT}"
		return 0
	else
		echo -e "$test_name ...... ${FAIL_PROMPT}"
		return 1
	fi
}

#-a表示普通数组
declare -a test_order=(
	"test_set"
	"test_get"
	"test_set again"
	"test_delete"
	"test_get_after_delete"
	"test_delete_after_delete"
)

#-A表示关联数组
declare -A test_func=(
	["test_set"]="test_set"
	["test_get"]="test_get"
	["test_set again"]="test_set"
	["test_delete"]="test_delete"
	["test_get_after_delete"]="test_get_after_delete"
	["test_delete_after_delete"]="test_delete_after_delete"
)

pass_count=0
fail_count=0

# NOTE: macos date does not support `date +%s%N`. Let's use the weird $TIMEFORMAT.
#
TIMEFORMAT="======================================
Run ${#test_order[@]} tests in %R seconds."

time {
	for testname in "${test_order[@]}"; do
		if run_test "${test_func[$testname]}" "$testname"; then
			((pass_count++))
		else
			((fail_count++))
		fi
	done
}

echo -e "\e[1;32m$pass_count\e[0m passed, \e[1;31m$fail_count\e[0m failed."