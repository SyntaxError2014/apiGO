#!/usr/bin/python

import sys


global parameters 
parameters = {}

global statusCode
statusCode = 200
global response 
response = ""

inParams = eval(sys.argv[2])
#inParams = eval("{'nume':['gheorghe']}")

for p in inParams:
    parameters[p] = inParams[p][0]

#print parameters['nume']

#d = dict(locals(), **globals())

exec(sys.argv[1])

print statusCode
print response