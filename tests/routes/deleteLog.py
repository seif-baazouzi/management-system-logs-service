import utils
import config
from testRestApi import testRoute, Test, DELETE

def testToken():
    logID = utils.createLog()
    res = testRoute(DELETE, f"{config.server}/api/v1/logs/{logID}")
    return res.equals({ "message": "invalid-token" })

def testDeleteLog():
    logID = utils.createLog()
    res = testRoute(DELETE, f"{config.server}/api/v1/logs/{logID}", headers={ "X-Token": config.token })
    return res.equals({ "message": "success"})


tests = [
    Test("Delete Log Invalid Token", testToken),
    Test("Delete Log", testDeleteLog),
]
