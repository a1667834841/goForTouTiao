import execjs
import requests
 
headers = {
    "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36",
}
 
url = "https://www.toutiao.com/api/pc/list/feed"
 
with open('../toutiao2.js', 'r', encoding='utf-8') as f:
    encrypt = f.read()
    _signature = execjs.compile(encrypt).call('getSignature')
 
params = {
    "channel_id": "0",
    "max_behot_time": "1664154636",
    "category": "pc_profile_recommend",
    "aid": "24",
    "app_name": "toutiao_web",
    "_signature": _signature
}
 
response = requests.get(url, params=params, verify=False)
print(response.text)