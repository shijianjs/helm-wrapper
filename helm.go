package main

import (
	"github.com/gin-gonic/gin"
	"os"

	"github.com/golang/glog"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/kube"
)

const XKubeToken = "X-KubeToken"
const XKubeApiServer = "X-KubeApiServer"

type KubeInformation struct {
	AimNamespace  string
	AimContext    string
	KubeToken     string
	KubeAPIServer string
}

func InitKubeInformation(namespace, context string, ginContext *gin.Context) *KubeInformation {
	return &KubeInformation{
		AimNamespace:  namespace,
		AimContext:    context,
		KubeToken:     ginContext.GetHeader(XKubeToken),
		KubeAPIServer: ginContext.GetHeader(XKubeApiServer),
	}
}

func actionConfigInit(kubeInfo *KubeInformation) (*action.Configuration, error) {
	actionConfig := new(action.Configuration)
	if kubeInfo.AimContext == "" {
		kubeInfo.AimContext = settings.KubeContext
	}
	clientConfig := kube.GetConfig(settings.KubeConfig, kubeInfo.AimContext, kubeInfo.AimNamespace)
	if kubeInfo.KubeToken != "" {
		clientConfig.BearerToken = &kubeInfo.KubeToken
		*clientConfig.Insecure = true
	}
	if kubeInfo.KubeAPIServer != "" {
		clientConfig.APIServer = &kubeInfo.KubeAPIServer
	}
	err := actionConfig.Init(clientConfig, kubeInfo.AimNamespace, os.Getenv("HELM_DRIVER"), glog.Infof)
	if err != nil {
		glog.Errorf("%+v", err)
		return nil, err
	}

	return actionConfig, nil
}
