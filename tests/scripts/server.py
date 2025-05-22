# Don't host on a public server!!
import requests
from flask import Flask, request

app = Flask(__name__)

@app.route('/ref_xss')
def xss():
    query = request.args.get('search', '')
    return f"<h1>Search results for: {query}</h1>"

@app.route('/dom_xss')
def dom_xss():
    return """
    <h1>Search results for:
    <script>document.write(new URLSearchParams(location.search).get("search"))</script>
    </h1>
    """

@app.route('/ssrf')
def ssrf():
    url = request.args.get('url', '')
    try: content = requests.get(url).text
    except: content = "Error fetching URL"
    return f"<h1>Content from {url}:</h1><pre>{content}</pre>"

@app.route('/rce')
def rce():
    cmd = request.args.get('cmd', '')
    try:
        import subprocess
        output = subprocess.check_output(cmd, shell=True, stderr=subprocess.STDOUT).decode()
    except Exception as e:
        output = str(e)
    return f"<h1>Command output:</h1><pre>{output}</pre>"

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')