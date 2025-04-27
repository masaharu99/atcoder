# 使用するベースイメージを指定
FROM golang:1.20

# 非rootユーザーの作成
RUN useradd -ms /bin/bash vscode
USER vscode
WORKDIR /workspace

# 必要な環境変数の設定
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

# 設定のキャッシュを利用してツールをインストール
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cweill/gotests/...@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
RUN go install github.com/cweill/gotests/...@latest
RUN go install golang.org/x/tools/gopls@latest
