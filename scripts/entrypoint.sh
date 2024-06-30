export http_proxy=${HTTP_PROXY}
export https_proxy=${HTTPS_PROXY}

pip install --no-cache-dir \
  httpx \
  jinja2

if [ -f /dependencies/apt-requirements.txt ]; then
  apt-get update
  apt-get install -y $(cat /dependencies/apt-requirements.txt)
fi

if [ -f /dependencies/python-requirements.txt ]; then
  pip install --no-cache-dir -r /dependencies/python-requirements.txt
fi

exec /main
