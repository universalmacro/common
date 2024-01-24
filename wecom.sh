url=$1
version=$2

org=universalmacro
generate_post_data()
{
  cat <<EOF
{
  "msgtype": "text",
  "text": {
    "content": "$org@$version SDK released\ngo get github.com/$org/@$version"
  }
}
EOF
}

curl -X POST $url -H 'Content-Type: application/json' --data "$(generate_post_data)"