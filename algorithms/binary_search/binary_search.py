

def binarySearchRecursive(arr, target, start, end):
    # find start index 
    # end index
    mid  = (start + end) // 2
    if arr[mid] == target:
    	return mid
    elif start == end:
    	return -1
    elif arr[mid] > target:
    	return binarySearchRecursive(arr, target, 0, mid-1)
    else:
    	return binarySearchRecursive(arr, target, mid+1, end)


def binarySearch(arr, target):

	left = 0
	right = len(arr) - 1

	while left <= right :
		mid = (left + right) // 2
		if arr[mid] == target :
			return mid
		elif arr[mid] > target:
			right = mid  - 1
		else:
			left = mid + 1

	return -1


def elementStatus(res, element):
	if res != -1:
		print(f"element {element} found at index {result}")
	else:
	    print(f"element {element} does not found in the data.")



if __name__ == '__main__':
	data = [1, 4, 9, 23, 43, 46, 78, 98]
	result = binarySearchRecursive(data, 79, 0, len(data))
	elementStatus(result, 79)

	result = binarySearchRecursive(data, 46, 0, len(data))
	elementStatus(result, 46)

	result = binarySearch(data, 4)
	elementStatus(result, 4)

	result = binarySearch(data, 99)
	elementStatus(result, 99)
