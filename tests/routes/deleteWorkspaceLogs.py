import utils
import config
from testRestApi import testRoute, Test, DELETE

def testToken():
    workspaceID = utils.createWorkspace()
    res = testRoute(DELETE, f"{config.server}/api/v1/logs/workspace/{workspaceID}")
    return res.equals({ "message": "invalid-token" })


def testNotExistingWorkspace():
    res = testRoute(DELETE, f"{config.server}/api/v1/logs/workspace/{utils.genUUID()}", headers={ "X-Token": config.token })
    return res.status == 400

def testDeleteWorkspaceLogs():
    workspaceID = utils.createWorkspace()
    res = testRoute(DELETE, f"{config.server}/api/v1/logs/workspace/{workspaceID}", headers={ "X-Token": config.token })
    return res.equals({ "message": "success"})


tests = [
    Test("Delete Workspace Logs Invalid Token", testToken),
    Test("Delete Workspace Logs Not Existing Workspace", testNotExistingWorkspace),
    Test("Delete Workspace Logs", testDeleteWorkspaceLogs),
]
