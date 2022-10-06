import random

import utils
import config
from testRestApi import testRoute, Test, PUT

def testToken():
    logID = utils.createLog()
    res = testRoute(PUT, f"{config.server}/api/v1/logs/{logID}")
    return res.equals({ "message": "invalid-token" })

def testEmptyFields():
    logID = utils.createLog()
    res = testRoute(PUT, f"{config.server}/api/v1/logs/{logID}", headers={ "X-Token": config.token })
    return res.equals({ "label": "Must not be empty" })

def testUpdateLog():
    logID = utils.createLog()
    body = {
        "label": utils.randomString(10),
        "description": utils.randomString(50),
        "value": random.randint(10, 50),
        "date": utils.today()
    }

    res = testRoute(PUT, f"{config.server}/api/v1/logs/{logID}", headers={ "X-Token": config.token }, body=body)
    return res.equals({ "message": "success"})


tests = [
    Test("Update Log Invalid Token", testToken),
    Test("Update Log Empty Fields", testEmptyFields),
    Test("Update Log", testUpdateLog),
]
