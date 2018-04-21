import os

'''
old_string = "file2344.txt"
to_remove = "0123456789"
table = str.maketrans("", "", to_remove)
new_string = old_string.translate(table)
print(new_string)
'''
'''
old_string = "file52.txt"
to_remove = "0123456789"
table = {ord(char): None for char in to_remove}
new_string = old_string.translate(table)
print(new_string)
#works int he commandline
'''

def rename_files():
	#(1) get file names from a folder
	file_list = os.listdir(r"C:\Users\g\Desktop\prank")
	print(file_list)

	#(2) for each file, rename filename
	for file_name in file_list:
		os.rename(file_name,)
rename_files()
