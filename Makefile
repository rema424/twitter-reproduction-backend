.PHONY: all

# 環境変数を読み込んで変数に代入する
# ただし、コマンドライン引数で値が渡された場合は、環境変数が定義されていてもコマンドライン引数の方を優先する
# e.g.) make xxxxxx project_id=xxxxxxxxxxxx
project_id := ${GOOGLE_PROJECT_ID}
version := ${GAE_VERSION}

# シェルコマンドの実行結果を代入する時は$でエスケープする
OAUTH2_ACCESS_TOKEN=$$(gcloud auth print-access-token 2> /dev/null)

#
# タスク定義
#

# DI
# e.g.) make wire
wire:
	cd module/twitter/main; wire;

# 認証
# e.g.) make login
login:
	gcloud auth activate-service-account --key-file key/deploy.json

# ローカルサーバーの起動
# e.g.) make dev
dev:
	dev_appserver.py module/twitter/main/app.yaml --support_datastore_emulator=False --go_debugging=True

# デバッガーのアタッチ
# e.g.) make debug
debug:
	@PID=$$(ps u | grep _go_ap[p] | head -1 | awk '{print $$2}'); dlv attach $$PID --headless --listen=127.0.0.1:2345 --api-version=2

# デプロイ
# e.g.) make update version=1
update:
	appcfg.py update -A ${project_id} -V ${version} --oauth2_access_token=${OAUTH2_ACCESS_TOKEN} module/twitter/main

# トラフィックを徐々に切り替える (500が発生しない 基本的にはこっち)
# e.g) make migrate version=1
migrate:
	appcfg.py migrate_traffic -A ${project_id} -V ${version} --oauth2_access_token=${OAUTH2_ACCESS_TOKEN} module/twitter/main

# トラフィックを一瞬で切り替える (一瞬500が発生する)
# e.g) make set version=1
set:
	appcfg.py set_default_version -A ${project_id} -V ${version} --oauth2_access_token=${OAUTH2_ACCESS_TOKEN} module/twitter/main

# e.g) make open version=1
open:
	gcloud app browse --version ${version}