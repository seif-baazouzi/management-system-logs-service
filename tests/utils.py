import uuid
import string
import random
import datetime

import config
from testRestApi import testRoute, POST

def randomString(length):
    return ''.join(random.choices(string.ascii_letters, k=length))

def genUUID():
    return str(uuid.uuid4())
    
def today():
    return datetime.datetime.utcnow().isoformat("T") + "Z"

def month():
    return today()[:7]

def createWorkspace():
    workspace = randomString(10)
    parentWorkspace = None

    body = { "workspace": workspace, "parentWorkspace": parentWorkspace }
    res = testRoute(POST, f"{config.workspacesService}/api/v1/workspaces", headers={ "X-Token": config.token }, body=body)
    
    workspaceID = False if "workspaceID" not in res.body else res.body["workspaceID"]

    return workspaceID

def createLog():
    workspaceID = createWorkspace()

    body = { "label": randomString(10), "value": random.randint(10, 50), "date": today() }
    res = testRoute(POST, f"{config.server}/api/v1/logs/{workspaceID}", headers={ "X-Token": config.token }, body=body)
    
    logID = False if "logID" not in res.body else res.body["logID"]

    return logID
