from os import getenv
from nut2 import PyNUTClient
from thingspeak import Channel

CHANNEL_ID = int(getenv("THINGSPEAK_CHANNEL_ID"))
API_KEY = getenv("THINGSPEAK_API_KEY")
UPS_NAME = getenv("NUT_UPS_NAME")

channel = Channel(CHANNEL_ID, api_key=API_KEY)
client = PyNUTClient()
voltage = float(client.get_var(UPS_NAME, 'input.voltage'))
print(voltage, channel.update({"field1": voltage}))
