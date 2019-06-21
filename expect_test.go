package aidi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestGetConsumerAccountExpectErr(t *testing.T) {
	type testCase struct {
		s1 string
		s2 string
		except bool
	}

	testCases := []testCase{
		testCase{
			s1: `[
    {
        "id": 1,
        "name": "adminName",
        "email": "zhang@megvii.com",
        "phone_number": "18566612312",
        "company_name": "adminCompany",
        "pushed": false
    },
    {
        "id": 2,
        "name": "Test Consumer 2",
        "email": "test2_consumer@gmail.com",
        "phone_number": "18566677974",
        "company_name": "Test2 Company",
        "pushed": false
    },
    {
        "id": 3,
        "name": "Test consumer 3",
        "email": "acv@gmail.com",
        "phone_number": "10086",
        "company_name": "c3",
        "pushed": false
    },
    {
        "id": 5,
        "name": "Test consumer 4",
        "email": "acv1@gmail.com",
        "phone_number": "10086",
        "company_name": "测试公司",
        "pushed": false
    },
    {
        "id": 6,
        "name": "adminName",
        "email": "zhang1@megvii.com",
        "phone_number": "18566612312",
        "company_name": "adminCompany",
        "pushed": false
    }
]`,
	s2 : `[
    {
        "id": 1,
        "name": "adminName",
        "email": "zhang@megvii.com",
        "phone_number": "18566612312",
        "company_name": "adminCompany",
        "pushed": false
    },
    {
        "id": 2,
        "name": "Test Consumer 2",
        "email": "test2_consumer@gmail.com",
        "phone_number": "18566677974",
        "company_name": "Test2 Company",
        "pushed": false
    },
    {
        "id": 3,
        "name": "Test consumer 3",
        "email": "acv@gmail.com",
        "phone_number": "10086",
        "company_name": "c3",
        "pushed": false
    },
    {
        "id": 5,
        "name": "Test consumer 4",
        "email": "acv1@gmail.com",
        "phone_number": "10086",
        "company_name": "测试公司",
        "pushed": false
    }
]`,
   except: true,
		},
		{
			s1: `[
    {
        "id": 1,
        "name": "adminName",
        "email": "zhang@megvii.com",
        "phone_number": "18566612312",
        "company_name": "adminCompany",
        "pushed": false
    },
    {
        "id": 2,
        "name": "Test Consumer 2",
        "email": "test2_consumer@gmail.com",
        "phone_number": "18566677974",
        "company_name": "Test2 Company",
        "pushed": false
    },
    {
        "id": 3,
        "name": "Test consumer 3",
        "email": "acv@gmail.com",
        "phone_number": "10086",
        "company_name": "c3",
        "pushed": false
    },
    {
        "id": 5,
        "name": "Test consumer 4",
        "email": "acv1@gmail.com",
        "phone_number": "10086",
        "company_name": "测试公司",
        "pushed": false
    },
    {
        "id": 6,
        "name": "adminName",
        "email": "zhang1@megvii.com",
        "phone_number": "18566612312",
        "company_name": "adminCompany",
        "pushed": false
    }
]`,
   s2: `[
    {
        "id": 1,
        "name": "adminName",
        "email": "zhang@megvii.com",
        "phone_number": "18566612312",
        "company_name": "adminCompany",
        "pushed": false
    },
    {
        "id": 2,
        "name": "Test Consumer 2",
        "email": "test2_consumer@gmail.com",
        "phone_number": "18566677974",
        "company_name": "Test2 Company",
        "pushed": false
    },
    {
        "id": 3,
        "name": "Test consumer 3",
        "email": "acv@gmail.com",
        "phone_number": "10086",
        "company_name": "c3",
        "pushed": false
    },
    {
        "id": 5,
        "email": "acv1@gmail.com",
        "phone_number": "10086",
        "company_name": "测试公司",
        "pushed": false
    }
]`,
			except: true,
		},
		{
			s1: `[
    {
        "id": 1,
        "name": "adminName",
        "email": "zhang@megvii.com",
        "phone_number": "18566612312",
        "company_name": "adminCompany",
        "pushed": false
    },
    {
        "id": 2,
        "name": "Test Consumer 2",
        "email": "test2_consumer@gmail.com",
        "phone_number": "18566677974",
        "company_name": "Test2 Company",
        "pushed": false
    }
]`,
			s2: `[
    {
        "id": 1,
        "name": "adminName1",
        "email": "zhang@megvii.com",
        "phone_number": "18566612312",
        "company_name": "adminCompany",
        "pushed": false
    }
]`,
			except: false,
		},

		{
			s1: `[
    {
        "id": 1,
        "name": "adminName",
        "email": "zhang@megvii.com",
        "phone_number": "18566612312",
        "company_name": "adminCompany",
        "pushed": false
    },
    {
        "id": 2,
        "name": "Test Consumer 2",
        "email": "test2_consumer@gmail.com",
        "phone_number": "18566677974",
        "company_name": "Test2 Company",
        "pushed": false
    }
]`,
			s2: `[
    {
        "id": 1,
        "name": "adminName",
        "email": "zhang@megvii.com",
        "company_name": "adminCompany",
        "pushed": false
    }
]`,
			except: true,
		},

		{
			s1: `[
    {
        "id": 1,
        "name": "adminName",
        "email": "zhang@megvii.com",
        "pushed": false
    }
]`,
			s2: `[
    {
        "id": 1,
        "name": "adminName",
        "email": "zhang@megvii.com",
        "company_name": "adminCompany",
        "pushed": false
    }
]`,
			except: false,
		},
		{
			s1: `
    {
        "id": 1,
        "name": "adminName",
        "email": "zhang@megvii.com",
        "pushed": false
    }
`,
			s2: `
    {
        "id": 1,
        "name": "adminName",
        "email": "zhang@megvii.com",
        "company_name": "adminCompany",
        "pushed": false
    }
`,
			except: false,
		},
		{
			s1: `
   {
       "id": 1,
       "name": "adminName"
   }
`,
			s2: `
   {
       "id": 1
   }
`,
			except: true,
		},

		{
			s1: `
   {
       "name": "adminName"
   }
`,
			s2: `
   {
       "id": 1
   }
`,
			except: false,
		},
	}

	for _, v := range testCases {
		c, err := containsJSON(v.s1, v.s2)
		assert.Nil(t, err)
		assert.Equal(t, c, v.except)
	}


}

