# Author:  Menglong TAN <tanmenglong@gmail.com>
# Date:    Tue May  2 20:47:23 2017
#
# Make Target:
# ------------
# The Makefile provides the following targets to make:
#   $ make           compile and link
#   $ make clean     clean objects and the executable file
#
#===========================================================================

all :
	thrift -r --out ../ --gen go ./nusadua.thrift

clean :
	rm -rf ../server
