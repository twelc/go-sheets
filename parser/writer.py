import gspread
import datetime
sheet_id = "141maOrpeeFsydVAWP-kIaziMCHn_fI8nQv0mFB78TVk"
gc = gspread.service_account('./roofsparser-addef44f7a5a.json')
sh = gc.open_by_key(sheet_id)

def create_tmp():
    try:
        sh.add_worksheet(title="tmp", rows=10000, cols=4)
    except gspread.exceptions.APIError:
        ws = sh.worksheet("tmp")
        sh.del_worksheet(ws)
        sh.add_worksheet(title="tmp", rows=10000, cols=4)

def save_data(data, index):
    ws = sh.worksheet("tmp")
    ws.update(f'A{str(index)}:C{str(index)}', data)

def accept_dump():
    try:
        sh.del_worksheet(sh.worksheet("default"))
    except:
        pass
    ws = sh.worksheet("tmp")
    ws.update_cell(4, 4, str(datetime.datetime.now()))
    ws.update_title("default")

