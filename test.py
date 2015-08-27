import mercury200 as m

try:
    print m.get_serial_number('266608','COM5',1)
except m.TimeoutError:
    print "suck because of timeout"






