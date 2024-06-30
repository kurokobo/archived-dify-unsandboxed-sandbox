pip install --no-cache-dir \
  httpx \
  jinja2

if [ -f /dependencies/python-requirements.txt ]; then
  pip install --no-cache-dir -r /dependencies/python-requirements.txt
fi

exec /main
