# Go 1.20 ベースイメージ
FROM golang:1.20

# ===== システム依存パッケージとツール類を root で導入 =====
# --- 抜粋（コメントを削除して改行を正しくエスケープ） ---
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
        git build-essential curl unzip \
        python3 python3-pip \
        nodejs npm \
        jq \
    && rm -rf /var/lib/apt/lists/*


# AtCoder CLI（acc）を npm 経由で導入
RUN npm install -g atcoder-cli           # :contentReference[oaicite:0]{index=0}

# online-judge-tools（oj）を pip 経由で導入
RUN pip3 install --no-cache-dir --break-system-packages online-judge-tools \
 && pip3 install --no-cache-dir --break-system-packages aclogin

# ===== 非 root ユーザー vscode を作成し、作業ディレクトリを設定 =====
RUN useradd -ms /bin/bash vscode
USER vscode
WORKDIR /workspace

# ===== Go 関連ツールの導入（ユーザー環境にインストール） =====
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.1 \
 && go install golang.org/x/tools/gopls@v0.15.3 \
 && go install github.com/cweill/gotests/...@latest \
 && go install github.com/ramya-rao-a/go-outline@latest \
 && go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest

RUN echo "alias ojt='oj t -c \"go run main.go\" -d ./tests'" >> /home/vscode/.bashrc