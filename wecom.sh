url=$1
packageName=$2
version=$3

generate_post_data()
{
  cat <<EOF
{
  "msgtype": "text",
  "text": {
    "content": "$packageName@$version SDK released"
  }
}
EOF
}

curl -X POST $url -H 'Content-Type: application/json' --data "$(generate_post_data)"