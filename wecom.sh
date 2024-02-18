url=$1
package=$2
version=$3

org=universalmacro
generate_post_data()
{
  cat <<EOF
{
  "msgtype": "text",
  "text": {
    "content": "$org@$version SDK released\ngo get github.com/$org/$package@$version"
  }
}
EOF
}

curl -X POST $url -H 'Content-Type: application/json' --data "$(generate_post_data)"