.PHONY: mount
mount:
	mount -t overlay overlay -olowerdir=lower1:lower2:lower3,upperdir=upper,workdir=work merged
	mount | grep lower1
