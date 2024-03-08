from ctypes import *

lib = cdll.LoadLibrary("./sheet_lib/lib.so")

cred = "./creds.json"
sheet_id = "..."
sheet_name = "default"


lib.AppendData.argtypes = [c_char_p, c_char_p, c_char_p, c_char_p, c_char_p, c_int, c_int]

lib.SaveLine.argtypes = [c_char_p, c_char_p, c_char_p, c_char_p, c_char_p, c_int, c_int]


def make_config(credentials:str, _sheet_id:str, _sheet_name:str):
    global cred
    global sheet_name
    global sheet_id
    cred = credentials
    sheet_id = _sheet_id
    sheet_name = _sheet_name

def append_data(obj_name:str, district:str, value:int, index:int):
    global cred
    global sheet_name
    global sheet_id
    if sheet_id == "...":
        raise AttributeError("you need to set the config, use make_config")
    lib.AppendData(cred.encode(), sheet_id.encode(), sheet_name.encode(), obj_name.encode(), district.encode(), value, index)


def save_line(obj_name:str, district:str, value:int, index:int):
    global cred
    global sheet_name
    global sheet_id
    if sheet_id == "...":
        raise AttributeError("you need to set the config, use make_config")
    lib.SaveLine(cred.encode(), sheet_id.encode(), sheet_name.encode(), obj_name.encode(), district.encode(), value, index)
