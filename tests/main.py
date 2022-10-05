from testRestApi import runTests

import routes.getLogs as getLogs

if __name__ == "__main__":
    runTests([
        *getLogs.tests,
    ])
