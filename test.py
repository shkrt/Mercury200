import mercury200
import serial

command = mercury200.make_command('266608','28')
print command

sp = serial.Serial("COM5")

sp.write(command)
response = sp.read(13)



#sp.write("")
sp.close()
