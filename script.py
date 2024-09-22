from os import getenv
from PyNUTClient.PyNUT import PyNUTClient
from thingspeak import Channel

CHANNEL_ID = int(getenv("THINGSPEAK_CHANNEL_ID"))
API_KEY = getenv("THINGSPEAK_API_KEY")
UPS_NAME = getenv("NUT_UPS_NAME")

channel = Channel(CHANNEL_ID, api_key=API_KEY)
client = PyNUTClient()
if client.CheckUPSAvailable(UPS_NAME):
    data = client.GetUPSVars(UPS_NAME)
    # use byte sequence instead of str as suggested by docs
    voltage = float(data[b'input.voltage'])
    print(voltage, channel.update({"field1": voltage}))
else:
    raise Exception("Failed to verify if ups was online")
