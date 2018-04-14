# Use
golangとcobraを使ってPC生活楽にしようぜ！ってことで習作です。  
コマンドリストを以下に記述していきます。  
todoもついでに書いていきます。  
# CommandList
- qqw gc  
Auto git to Add + Commit + Push  
- qqw toot ...  
マストドンにトゥートします。  
- qqw tl
マストドンのローカルタイムラインを見つめます。  
連合は作者が使ってないので見れませんというクズっぷりです。  
- qqw tweet
とりあえず、Twitterにツイートはできるようになりました。  

# Develop Platform
Linux 4.13  
Xubuntu 16.04 i686  
  
# Config
- Mastodonの設定ファイルを外部ファイル化しました。
example.qqw.ymlの中身を書き換えて、.cdg.jsonにして$HOMEだけを見てるのでそこにおいてください。  
- ClientIDとClientSecretの取得方法ですが、web版のマストドンを開き
1. 設定
2. 開発
3. 新規アプリ
4. アプリの名前
5. アプリのウェブサイト
6. 送信
  
アプリ名のページを開き  
ここからクライアントキーとクライアントシークレットがコピーできます  
  
なんか、[dev.twitter.com](https://developer.twitter.com/)から直でアプリ管理画面いけなくなってません？  
[apps.twitter.com](https://apps.twitter.com/)をURL直打ちするとは思わなかった  

# 一応…
qqw tl を使うとカラー指定と、DisplayName(Username)をしてるので、  
アイコンは見えませんが、IDと表示名見えてたら捕捉できるよね？  

