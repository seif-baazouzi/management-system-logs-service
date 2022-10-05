from testRestApi import runTests

import routes.getLogs as getLogs
import routes.createLog as createLog

if __name__ == "__main__":
    runTests([
        *getLogs.tests,
        *createLog.tests,
    ])
