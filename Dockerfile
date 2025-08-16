FROM golang:1.23

# # 後述のci.ymlから渡される環境変数
# ARG ARTIFACT_PATH

# COPY ${ARTIFACT_PATH} /main

CMD ["/main"]