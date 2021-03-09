project('RoboDep', 'c', default_options : ['warning_level=3', 'c_std=c99'])

add_global_arguments(language : 'c')
c = meson.get_compiler('c')

#Sample built-in libs
#mathl = c.find_library('m', required : true)
#runtimel = c.find_library('rt', required : true)
#dlib = c.find_library('dl', required : true)

#Sample dependencies
#depz = [dependency('raylib'), dependency('GL'), dependency('X11'), dependency('threads')]

#include dirs
#libz = include_directories('lib')

#main source
srcs = ['src/main.c']

executable('main', srcs)