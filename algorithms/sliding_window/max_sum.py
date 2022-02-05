import sys


def maxSum(arr, k):
	resp  = {"num": -2  *  sys.maxsize -1, "idx": 0}

	for i in range(0, len(arr) + 1 - k):
		curr_sum = sum(arr[i:i+k])

		if curr_sum > resp["num"]:
			resp["num"] = curr_sum
			resp["idx"] = i
		

	return resp


def slidingWindow(arr, k):
	resp  = {"num": sum(arr[0:k]), "idx": 0}
	for i in range(0, len(arr) - k):
		new_sum = resp["num"] - arr[i] + arr[i+k]

		if new_sum  > resp["num"]:
			resp["idx"] = i + 1
		resp["num"] = new_sum

	return resp

if __name__ == '__main__':
	data = [21, 91, 13, 41, 82, 49, 28, 30, 49, 29, 19, 83, 66, 32]
	print(maxSum(data, 3))
	print(slidingWindow(data, 3))