package testlink

// TestLinkMethod
const (
	TestLinkMethodAssignTestCaseExecutionTask               = "tl.assignTestCaseExecutionTask"
	TestLinkMethodCreateTestProject                         = "tl.createTestProject"
	TestLinkMethodGetTestProjectByName                      = "tl.getTestProjectByName"
	TestLinkMethodCheckDevKey                               = "tl.checkDevKey"
	TestLinkMethodGetProjects                               = "tl.getProjects"
	TestLinkMethodCreateTestPlan                            = "tl.createTestPlan"
	TestLinkMethodGetTestPlanByName                         = "tl.getTestPlanByName"
	TestLinkMethodGetProjectTestPlans                       = "tl.getProjectTestPlans"
	TestLinkMethodGetTestPlanPlatforms                      = "tl.getTestPlanPlatforms"
	TestLinkMethodCreateTestCase                            = "tl.createTestCase"
	TestLinkMethodUpdateTestCase                            = "tl.updateTestCase"
	TestLinkMethodCreateTestCaseSteps                       = "tl.createTestCaseSteps"
	TestLinkMethodDeleteTestCaseSteps                       = "tl.deleteTestCaseSteps"
	TestLinkMethodCreateTestSuite                           = "tl.createTestSuite"
	TestLinkMethodAddTestCaseToTestPlan                     = "tl.addTestCaseToTestPlan"
	TestLinkMethodDoesUserExist                             = "tl.doesUserExist"
	TestLinkMethodGetUserByLogin                            = "tl.getUserByLogin"
	TestLinkMethodSayHello                                  = "tl.sayHello"
	TestLinkMethodAbout                                     = "tl.about"
	TestLinkMethodGetTestSuiteById                          = "tl.getTestSuiteByID"
	TestLinkMethodSetTestMode                               = "tl.setTestMode"
	TestLinkMethodRepeat                                    = "tl.repeat"
	TestLinkMethodGetTestCase                               = "tl.getTestCase"
	TestLinkMethodGetTestCasesForTestSuite                  = "tl.getTestCasesForTestSuite"
	TestLinkMethodCreateBuild                               = "tl.createBuild"
	TestLinkMethodGetTestCasesForTestPlan                   = "tl.getTestCasesForTestPlan"
	TestLinkMethodGetTestCaseIdByName                       = "tl.getTestCaseIDByName"
	TestLinkMethodGetTestSuitesForTestPlan                  = "tl.getTestSuitesForTestPlan"
	TestLinkMethodUploadAttachment                          = "tl.uploadAttachment"
	TestLinkMethodUploadTestCaseAttachment                  = "tl.uploadTestCaseAttachment"
	TestLinkMethodUploadTestSuiteAttachment                 = "tl.uploadTestSuiteAttachment"
	TestLinkMethodUploadTestProjectAttachment               = "tl.uploadTestProjectAttachment"
	TestLinkMethodUploadRequirementAttachment               = "tl.uploadRequirementAttachment"
	TestLinkMethodUploadRequirementSpecificationAttachment  = "tl.uploadRequirementSpecificationAttachment"
	TestLinkMethodGetTestCaseAttachments                    = "tl.getTestCaseAttachments"
	TestLinkMethodGetTestSuiteAttachments                   = "tl.getTestSuiteAttachments"
	TestLinkMethodUploadExecutionAttachment                 = "tl.uploadExecutionAttachment"
	TestLinkMethodDeleteExecution                           = "tl.deleteExecution"
	TestLinkMethodGetFullPath                               = "tl.getFullPath"
	TestLinkMethodAssignRequirements                        = "tl.assignRequirements"
	TestLinkMethodGetTestSuitesForTestSuite                 = "tl.getTestSuitesForTestSuite"
	TestLinkMethodGetFirstLevelTestSuitesForTestProject     = "tl.getFirstLevelTestSuitesForTestProject"
	TestLinkMethodReportTcResult                            = "tl.reportTCResult"
	TestLinkMethodGetLastExecutionResult                    = "tl.getLastExecutionResult"
	TestLinkMethodGetBuildsForTestPlan                      = "tl.getBuildsForTestPlan"
	TestLinkMethodGetLatestBuildForTestPlan                 = "tl.getLatestBuildForTestPlan"
	TestLinkMethodGetTestCaseKeywords                       = "tl.getTestCaseKeywords"
	TestLinkMethodGetTestCaseCustomFieldDesignValue         = "tl.getTestCaseCustomFieldDesignValue"
	TestLinkMethodGetTestCaseCustomFieldTestPlanDesignValue = "tl.getTestCaseCustomFieldTestPlanDesignValue"
	TestLinkMethodGetTestCaseCustomFieldExecutionValue      = "tl.getTestCaseCustomFieldExecutionValue"
	TestLinkMethodGetTestPlanCustomFieldDesignValue         = "tl.getTestPlanCustomFieldDesignValue"
	TestLinkMethodGetTotalsForTestPlan                      = "tl.getTotalsForTestPlan"
	TestLinkMethodGetExecCountersByBuild                    = "tl.getExecCountersByBuild"
	TestLinkMethodUpdateTestCaseCustomFieldValue            = "tl.updateTestCaseCustomFieldDesignValue"
	TestLinkMethodSetTestCaseExecutionType                  = "tl.setTestCaseExecutionType"
	TestLinkMethodGetProjectPlatforms                       = "tl.getProjectPlatforms"
	TestLinkMethodRemovePlatformFromTestPlan                = "tl.removePlatformFromTestPlan"
	TestLinkMethodAddPlatformToTestPlan                     = "tl.addPlatformToTestPlan"
	TestLinkMethodUpdateBuildCustomFields                   = "tl.updateBuildCustomFieldsValues"
	TestLinkMethodAddTestCaseKeyWords                       = "tl.addTestCaseKeywords"
	TestLinkMethodGetTestCaseBugs                           = "tl.getTestCaseBugs"
)

// TestLinkParams
const (
	TestLinkParamActionOnDuplicatedName     = "actiononduplicatedname"
	TestLinkParamActions                    = "actions"
	TestLinkParamAction                     = "action"
	TestLinkParamActive                     = "active"
	TestLinkParamAssignedTo                 = "assignedto"
	TestLinkParamAuthorLogin                = "authorlogin"
	TestLinkParamBugId                      = "bugid"
	TestLinkParamBuildId                    = "buildid"
	TestLinkParamBuildName                  = "buildname"
	TestLinkParamBuildNotes                 = "buildnotes"
	TestLinkParamCheckDuplicatedName        = "checkduplicatedname"
	TestLinkParamContent                    = "content"
	TestLinkParamCustomFieldName            = "customfieldname"
	TestLinkParamCustomFields               = "customfields"
	TestLinkParamDeep                       = "deep"
	TestLinkParamDescription                = "description"
	TestLinkParamDetails                    = "details"
	TestLinkParamDevKey                     = "devKey"
	TestLinkParamEnableAutomation           = "automationEnabled"
	TestLinkParamEnableInventory            = "inventoryEnabled "
	TestLinkParamEnableRequirements         = "requirementsEnabled"
	TestLinkParamEnableTestPriority         = "testPriorityEnabled"
	TestLinkParamEstimatedExecutionDuration = "estimated_execution_duration"
	TestLinkParamExecuted                   = "executed"
	TestLinkParamExecuteStatus              = "executestatus"
	TestLinkParamExecution                  = "execution"
	TestLinkParamExecutionId                = "executionid"
	TestLinkParamExecutionDuration          = "execduration"
	TestLinkParamExecutionType              = "executiontype"
	TestLinkParamStepExecutionType          = "execution_type"
	TestLinkParamExpectedResults            = "expected_results"
	TestLinkParamFileName                   = "filename"
	TestLinkParamFileType                   = "filetype"
	TestLinkParamFkId                       = "fkid"
	TestLinkParamFkTable                    = "fktable"
	TestLinkParamGetStepInfo                = "getstepsinfo"
	TestLinkParamGuess                      = "guess"
	TestLinkParamImportance                 = "importance"
	TestLinkParamInternalId                 = "internalId"
	TestLinkParamKeywordId                  = "keywordid"
	TestLinkParamKeywords                   = "keywords"
	TestLinkParamName                       = "name"
	TestLinkParamNodeId                     = "nodeid"
	TestLinkParamNotes                      = "notes"
	TestLinkParamOptions                    = "options"
	TestLinkParamOrder                      = "order"
	TestLinkParamOverwrite                  = "overwrite"
	TestLinkParamParentId                   = "parentid"
	TestLinkParamPlatformId                 = "platformid"
	TestLinkParamPlatformName               = "platformname"
	TestLinkParamPreconditions              = "preconditions"
	TestLinkParamPublic                     = "public"
	TestLinkParamRequirementSpecificationId = "reqspecid"
	TestLinkParamRequirementId              = "requirementid"
	TestLinkParamRequirements               = "requirements"
	TestLinkParamResult                     = "result"
	TestLinkParamSummary                    = "summary"
	TestLinkParamStatus                     = "status"
	TestLinkParamStepNumber                 = "step_number"
	TestLinkParamSteps                      = "steps"
	TestLinkParamStr                        = "str"
	TestLinkParamTestProjectName            = "testprojectname"
	TestLinkParamTestCaseExternalId         = "testcaseexternalid"
	TestLinkParamTestCaseId                 = "testcaseid"
	TestLinkParamTestCaseName               = "testcasename"
	TestLinkParamTestCasePathName           = "testcasepathname"
	TestLinkParamTestCasePrefix             = "testcaseprefix"
	TestLinkParamTestMode                   = "testmode"
	TestLinkParamTestPlanId                 = "testplanid"
	TestLinkParamTestPlanName               = "testplanname"
	TestLinkParamTestProjectId              = "testprojectid"
	TestLinkParamTestSuiteId                = "testsuiteid"
	TestLinkParamTestSuiteId2               = "testsuite_id"
	TestLinkParamTestSuiteName              = "testsuitename"
	TestLinkParamTimestamp                  = "timestamp"
	TestLinkParamTitle                      = "title"
	TestLinkParamUrgency                    = "urgency"
	TestLinkParamUser                       = "user"
	TestLinkParamVersion                    = "version"
)