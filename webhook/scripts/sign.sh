#!/bin/bash

# Initialize our own variables:
message=""
secret=""
decode=false

# Parse the command line arguments
while getopts ":m:s:d" opt; do
  case $opt in
    m)
      message="$OPTARG"
      ;;
    s)
      secret="$OPTARG"
      ;;
    d)
      decode=true
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      ;;
  esac
done

if $decode ; then
  secret=$(echo "$secret" | base64 --decode)
fi

# We are using json.Marshal to generate the json, and it does html escaping to strings so we have to do that here also
# ref.: https://pkg.go.dev/encoding/json#Marshal
# Note however that we only need this because the script takes the message as a string
# In your program you'll have access to the message as a byte array and you should use the byte array directly
htmlEscapedMessage=$(echo -n "$message" | sed 's/</\\u003c/g' | sed 's/>/\\u003e/g' | sed 's/&/\\u0026/g')

# echo "Message: "
# echo "$message"
# echo "Escaped string: "
# echo "$htmlEscapedMessage"
# echo "Secret: $secret"

signature=$(echo -n "$htmlEscapedMessage" | openssl dgst -sha256 -hmac $secret)

echo "Signature: $signature"