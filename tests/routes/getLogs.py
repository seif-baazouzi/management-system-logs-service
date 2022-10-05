import utils
import config
from testRestApi import testRoute, Test, GET

def testToken():
    workspaceID = utils.createWorkspace()
    res = testRoute(GET, f"{config.server}/api/v1/logs/{workspaceID}/month/{utils.month()}")
    return res.equals({ "message": "invalid-token" })

def testGetMonthTodos():
    workspaceID = utils.createWorkspace()
    res = testRoute(GET, f"{config.server}/api/v1/logs/{workspaceID}/month/{utils.month()}", headers={ "X-Token": config.token })
    return res.hasKeys("logs")

tests = [
    Test("Get Logs Invalid Token", testToken),
    Test("Get Logs", testGetMonthTodos),
]
