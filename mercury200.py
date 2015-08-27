import serial
import array
from PyCRC.CRC16 import CRC16


# makes list from crc bytes
def hex_to_hexarr(num):
    res = []
    while num > 0:
        byte = num & 0xff
        res.append(int(byte))
        num >>= 8
    return res


# makes list from counter's serial number
def string_to_list(num):
    intnum = int(num);
    hexnum = "{0:0{1}x}".format(intnum, 6)
    return [hexnum[i:i + 2] for i in range(0, len(hexnum), 2)]


# makes command from given serial number, command code
def make_command(serial, code):
    request = ['0']
    x = string_to_list(serial)  # '266608'
    next = request + x + [code]  # '28'
    lst = map(lambda x: int(x, 16), next)
    arr = array.array('B', lst).tostring()
    crc = CRC16(True).calculate(arr)
    return lst + hex_to_hexarr(crc)


def check_crc(response):
    crcbytes = map(ord, response[-2:])
    resp = response[0:-2]
    crc = hex_to_hexarr(CRC16(True).calculate(resp))

    if crc == crcbytes:
        return True
    return False


def print_bytes(bt):
    l = []
    for b in bt:
        l.append(b)
    print l


# COUNTER COMMANDS
def get_version(net_number, portname,_timeout):
    command = make_command(net_number, '28')
    sp = serial.Serial(portname,timeout=_timeout)
    sp.write(command)
    response = sp.read(13)
    sp.close()
    if len(response)==0:
        return 'NO ANSWER FROM COUNTER'
    if check_crc(response):
        verbytes = map(ord, response[5:7])
        verchars = map(str, verbytes)
        return '.'.join(verchars)
