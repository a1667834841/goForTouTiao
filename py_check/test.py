import requests

url = "https://mp.toutiao.com/mp/agw/article/wtt"

payload = "{\"content\":\"自动上传again！！\",\"image_list\":[],\"is_fans_article\":2,\"pre_upload\":1,\"welfare_card\":\"\"}"
headers = {
  'authority': 'mp.toutiao.com',
  'accept': 'application/json, text/plain, */*',
  'accept-language': 'zh-CN,zh;q=0.9,zh-TW;q=0.8,en-US;q=0.7,en;q=0.6',
  'cache-control': 'no-cache',
  'content-type': 'application/json; charset=UTF-8',
  'cookie': 'passport_csrf_token=8ff93566da8c76beca619a45e601141f; passport_csrf_token_default=8ff93566da8c76beca619a45e601141f; _ba=BA0.2-20220212-5110e-Ij0PHzQcGJTEuP4DyvRi; MONITOR_WEB_ID=2f5a5fa8-ef0b-4f43-9e67-9bef5f0fb52b; s_v_web_id=verify_la0ko8rw_SHOXYXQp_jdzg_4qqu_8U4e_DNjOimpXmdHN; tt_webid=7063765302613542414; _tea_utm_cache_2018=undefined; d_ticket=de0573bcbe96caee29ca3a581f81b6a04284a; sso_auth_status=58bed36b755213000d05d155b6176388; sso_auth_status_ss=58bed36b755213000d05d155b6176388; csrf_session_id=03017f3bfdf779674aeea2b3a07b0e21; _ga=GA1.2.657630763.1667629263; _gid=GA1.2.646170897.1667629263; msToken=mDIWmtMzAQNrfHjhidstOnY6tTEdytCV3MQZ1KznFgkVX9O9LhzSn0zY7lbPRLmiC4WYddX2bEXs3Dqgu7FMG7vdlFIKBoqMoesm2mYU1S4=; ttwid=1%7CqhXMuMk8ouSSbPWpXsxu2bj0fDLLEQRPj5Jn-dPZ0t4%7C1667647330%7C18545fd0eee9d5c34d38ca6eef8a153bfe77ab3ef38c27ae7e33ea9c1389c7b9; n_mh=vbG5bigRXNObdiel2kfrZiLDxUcCBOeukm7Ilg2_rYY; sso_uid_tt=1b20ae34f81f790a449240c303223c62; sso_uid_tt_ss=1b20ae34f81f790a449240c303223c62; toutiao_sso_user=46138964b9ebe6fee730820ff9bcb626; toutiao_sso_user_ss=46138964b9ebe6fee730820ff9bcb626; sid_ucp_sso_v1=1.0.0-KGEzMDFiOGUwZWUzZGI5ODAzOTA2NjFjZjA5YmRiNTU1NzliYzhhZjMKHQis4-yAwQEQ_46ZmwYYzwkgDDDQorPABTgGQPQHGgJobCIgNDYxMzg5NjRiOWViZTZmZWU3MzA4MjBmZjliY2I2MjY; ssid_ucp_sso_v1=1.0.0-KGEzMDFiOGUwZWUzZGI5ODAzOTA2NjFjZjA5YmRiNTU1NzliYzhhZjMKHQis4-yAwQEQ_46ZmwYYzwkgDDDQorPABTgGQPQHGgJobCIgNDYxMzg5NjRiOWViZTZmZWU3MzA4MjBmZjliY2I2MjY; odin_tt=bc88cdd448f761d98693fc9f660523c96ec535a8272d7a9d524d041390dbda8218a58fd2b3cb59da837cf4ecad258035; passport_auth_status=5f5cf3a8c2d2ff1d83a14abb21d4861d%2C0c26fe97f1f6555c16e15defb1a44b7e; passport_auth_status_ss=5f5cf3a8c2d2ff1d83a14abb21d4861d%2C0c26fe97f1f6555c16e15defb1a44b7e; sid_guard=c45b99043a9f714da4452201fd996426%7C1667647361%7C5183998%7CWed%2C+04-Jan-2023+11%3A22%3A39+GMT; uid_tt=cebb506096aa3bf61edf9da08c6a47bb; uid_tt_ss=cebb506096aa3bf61edf9da08c6a47bb; sid_tt=c45b99043a9f714da4452201fd996426; sessionid=c45b99043a9f714da4452201fd996426; sessionid_ss=c45b99043a9f714da4452201fd996426; sid_ucp_v1=1.0.0-KDNiNjY2NDc0NmU5MzA2MGE2NjNjMDIyZmY4YjlhZjQ1MWFiYmQ2NGYKFwis4-yAwQEQgY-ZmwYYzwkgDDgGQPQHGgJobCIgYzQ1Yjk5MDQzYTlmNzE0ZGE0NDUyMjAxZmQ5OTY0MjY; ssid_ucp_v1=1.0.0-KDNiNjY2NDc0NmU5MzA2MGE2NjNjMDIyZmY4YjlhZjQ1MWFiYmQ2NGYKFwis4-yAwQEQgY-ZmwYYzwkgDDgGQPQHGgJobCIgYzQ1Yjk5MDQzYTlmNzE0ZGE0NDUyMjAxZmQ5OTY0MjY; gftoken=YzQ1Yjk5MDQzYXwxNjY3NjQ3MzY1NzN8fDAGBgYGBgY; tt_scid=BcYPKjwoRi5PKA.vLaXcGdTgN0zSQtsg5vtYz50vQ3G33gq.5GUlq9vaYlll9EoKfafc',
  'origin': 'https://mp.toutiao.com',
  'pragma': 'no-cache',
  'referer': 'https://mp.toutiao.com/profile_v4/weitoutiao/publish',
  'sec-ch-ua': '"Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"',
  'sec-ch-ua-mobile': '?0',
  'sec-ch-ua-platform': '"Windows"',
  'sec-fetch-dest': 'empty',
  'sec-fetch-mode': 'cors',
  'sec-fetch-site': 'same-origin',
  'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36'
}

response = requests.request("POST", url, headers=headers, data=payload.encode())

print(response.text)
