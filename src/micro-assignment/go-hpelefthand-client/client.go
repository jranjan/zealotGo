package boilerplate

type hpeLefthandClient struct {
    api_url string
    username string
    password string
    cached_session_key string
}

type vsaNodeInfo struct {
    ip string
}


func (* hpeLefthandClient) Login() {
    return nil
}

func (* hpeLefthandClient) Logout() {
    return nil
}

func (* hpeLefthandClient) GetClusterDetails(clusterName string) {
    return nil
}

func (* hpeLefthandClient) FlexUp(cname string, v vsaNodeInfo) bool {
    return true
}

func (* hpeLefthandClient) FlexDown(cname string, v vsaNodeInfo) bool {
    return true
}
