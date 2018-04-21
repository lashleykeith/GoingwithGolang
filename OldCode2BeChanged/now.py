import os
def rename_files():
	#(1) get files names from a folder
	file_list = os.listdir(r"C:\\Users\\g\\Desktop\\Fun\\udacity\\foundationsofpython\\07finalrenamingapp\\prank")
	print(file_list)
	saved_path = os.getcwd()
	print (saved_path)
	os.chdir(r"C:\\Users\\g\\Desktop\\Fun\\udacity\\foundationsofpython\\07finalrenamingapp\\prank")

	#2 rename all files names in folder
	for file_name in file_list:
		print ("Old Name - " +file_name)
		print ("New Name - " +file_name.strip("0123456789"))
		os.rename(file_name,file_name.strip("0123456789"))
	os.chdir(saved_path)

rename_files()