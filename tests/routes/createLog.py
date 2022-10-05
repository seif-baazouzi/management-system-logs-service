import random

import utils
import config
from testRestApi import testRoute, Test, POST

def testToken():
    workspaceID = utils.createWorkspace()
    res = testRoute(POST, f"{config.server}/api/v1/logs/{workspaceID}")
    return res.equals({ "message": "invalid-token" })

def testEmptyFields():
    workspaceID = utils.createWorkspace()
    res = testRoute(POST, f"{config.server}/api/v1/logs/{workspaceID}", headers={ "X-Token": config.token })
    return res.equals({ "label": "Must not be empty" })

def testNotExistingWorkspace():
    res = testRoute(POST, f"{config.server}/api/v1/logs/{utils.genUUID()}", headers={ "X-Token": config.token })
    return res.status == 400

def testCreateLog():
    workspaceID = utils.createWorkspace()
    body = {
        "label": utils.randomString(10),
        "description": utils.randomString(50),
        "value": random.randint(10, 50),
        "date": utils.today()
    }

    res = testRoute(POST, f"{config.server}/api/v1/logs/{workspaceID}", headers={ "X-Token": config.token }, body=body)
    return res.equals({ "message": "success"})


tests = [
    Test("Create Log Invalid Token", testToken),
    Test("Create Log Empty Fields", testEmptyFields),
    Test("Create Log Not Existing Workspace", testNotExistingWorkspace),
    Test("Create Log", testCreateLog),
]
