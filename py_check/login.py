import requests
login_api = 'https://mp.toutiao.com/mp/agw/media/user_login_status_api'
headers = {
        'user-agent' : 'Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36',
        'referer' : 'https://mp.toutiao.com/profile_v4/index',
        'cookie' : 'passport_csrf_token=8ff93566da8c76beca619a45e601141f; passport_csrf_token_default=8ff93566da8c76beca619a45e601141f; _ba=BA0.2-20220212-5110e-Ij0PHzQcGJTEuP4DyvRi; MONITOR_WEB_ID=2f5a5fa8-ef0b-4f43-9e67-9bef5f0fb52b; s_v_web_id=verify_la0ko8rw_SHOXYXQp_jdzg_4qqu_8U4e_DNjOimpXmdHN; tt_webid=7063765302613542414; _tea_utm_cache_2018=undefined; d_ticket=de0573bcbe96caee29ca3a581f81b6a04284a; n_mh=keFwevYiSV6QXtd6LOzobEGmiZ-0aFauRm9Fl_798C0; sso_auth_status=58bed36b755213000d05d155b6176388; sso_auth_status_ss=58bed36b755213000d05d155b6176388; sso_uid_tt=7ea9153cb3c00c0ea61fc76d11e6f7dd; sso_uid_tt_ss=7ea9153cb3c00c0ea61fc76d11e6f7dd; toutiao_sso_user=44b29d31a2543bcf6678c07cbbf2c45d; toutiao_sso_user_ss=44b29d31a2543bcf6678c07cbbf2c45d; sid_ucp_sso_v1=1.0.0-KDE3NjU0YWM2OWM3MWIwN2RiYjNmNGE1MmQyOTk4NDhjNjY1MGYzY2EKHwjDjZDD8oyAAxCeuZebBhjPCSAMMLaLzIMGOAJA8QcaAmhsIiA0NGIyOWQzMWEyNTQzYmNmNjY3OGMwN2NiYmYyYzQ1ZA; ssid_ucp_sso_v1=1.0.0-KDE3NjU0YWM2OWM3MWIwN2RiYjNmNGE1MmQyOTk4NDhjNjY1MGYzY2EKHwjDjZDD8oyAAxCeuZebBhjPCSAMMLaLzIMGOAJA8QcaAmhsIiA0NGIyOWQzMWEyNTQzYmNmNjY3OGMwN2NiYmYyYzQ1ZA; passport_auth_status=0c26fe97f1f6555c16e15defb1a44b7e%2Ceaec3b8b7ed83a114b1085dc436e6bff; passport_auth_status_ss=0c26fe97f1f6555c16e15defb1a44b7e%2Ceaec3b8b7ed83a114b1085dc436e6bff; sid_guard=fd4df6f8d8309754898a5b79b6bb336c%7C1667619998%7C5184000%7CWed%2C+04-Jan-2023+03%3A46%3A38+GMT; uid_tt=9c836bed85a4b5a461175a1451c4de4e; uid_tt_ss=9c836bed85a4b5a461175a1451c4de4e; sid_tt=fd4df6f8d8309754898a5b79b6bb336c; sessionid=fd4df6f8d8309754898a5b79b6bb336c; sessionid_ss=fd4df6f8d8309754898a5b79b6bb336c; sid_ucp_v1=1.0.0-KDM1ODcxNGEwMDM4NTEyM2ZjMGM4ODk0YTVlZjAxZTM3ZDU1NjY1N2EKGQjDjZDD8oyAAxCeuZebBhjPCSAMOAJA8QcaAmhsIiBmZDRkZjZmOGQ4MzA5NzU0ODk4YTViNzliNmJiMzM2Yw; ssid_ucp_v1=1.0.0-KDM1ODcxNGEwMDM4NTEyM2ZjMGM4ODk0YTVlZjAxZTM3ZDU1NjY1N2EKGQjDjZDD8oyAAxCeuZebBhjPCSAMOAJA8QcaAmhsIiBmZDRkZjZmOGQ4MzA5NzU0ODk4YTViNzliNmJiMzM2Yw; csrf_session_id=03017f3bfdf779674aeea2b3a07b0e21; gftoken=ZmQ0ZGY2ZjhkOHwxNjY3NjIwMjgwOTV8fDAGBgYGBgY; odin_tt=2f954541366b3ae5ff29c180ed61087c63b1b8a33fff5a2103d0de9642f7ef2fc493eaa417c1c3c546ac44825109a9aa; ttwid=1%7CqhXMuMk8ouSSbPWpXsxu2bj0fDLLEQRPj5Jn-dPZ0t4%7C1667621932%7C632ffcacb69c5b5548a70440a0db858d74aa36c7485686ca9f026a59c5085dcd; tt_scid=cnmt6.uabgyvs8-PuN9gB9-uLmbkrz1FwPo4dx9XJ32CRsyXncOtQOhbBlriVnKbb685; _ga=GA1.2.657630763.1667629263; _gid=GA1.2.646170897.1667629263; msToken=mDIWmtMzAQNrfHjhidstOnY6tTEdytCV3MQZ1KznFgkVX9O9LhzSn0zY7lbPRLmiC4WYddX2bEXs3Dqgu7FMG7vdlFIKBoqMoesm2mYU1S4='
        }


def login():
    print(requests.get(login_api, headers=headers).json())


login()