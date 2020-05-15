import requests


"""
JSON RPC client using Python
"""


def rpc_call():
    url = 'http://localhost:2334/rpc'
    r = requests.post(url, json={
        'id': 1,
        'method': 'Arith.Multiply',
        'params': [{'A': 7, 'B': 3}]
    })
    print(r.text)

    r = requests.post(url, json={
        'id': 2,
        'method': 'Arith.Divide',
        'params': [{'A': 9, 'B': 2}]
    })
    print(r.text)


if __name__ == '__main__':
    rpc_call()
