# Author:  Menglong TAN <tanmenglong@gmail.com>
# Date:    Tue May  9 00:48:04 2017
#
# Make Target:
# ------------
# The Makefile provides the following targets to make:
#   $ make           compile and link
#   $ make clean     clean objects and the executable file
#
#===========================================================================

all :
	make -C thrift
	go build

clean :
	make clean -C thrift
