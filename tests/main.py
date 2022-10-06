from testRestApi import runTests

import routes.getLogs as getLogs
import routes.createLog as createLog
import routes.updateLog as updateLog
import routes.deleteLog as deleteLog

if __name__ == "__main__":
    runTests([
        *getLogs.tests,
        *createLog.tests,
        *updateLog.tests,
        *deleteLog.tests,
    ])
