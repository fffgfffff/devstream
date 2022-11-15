package general

import (
	"github.com/devstream-io/devstream/internal/pkg/configmanager"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/ci"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/ci/step"
	"github.com/devstream-io/devstream/pkg/util/log"
	"github.com/devstream-io/devstream/pkg/util/scm"
)

func preConfigGithub(options configmanager.RawOptions) error {
	opts, err := ci.NewCIOptions(options)
	if err != nil {
		return err
	}

	stepConfigs := step.ExtractValidStepConfig(opts.Pipeline)
	githubClient, err := scm.NewClientWithAuth(opts.ProjectRepo)
	if err != nil {
		log.Debugf("init github client failed: %+v", err)
		return err
	}
	for _, c := range stepConfigs {
		err := c.ConfigSCM(githubClient)
		if err != nil {
			log.Debugf("githubaction config github failed: %+v", err)
			return err
		}
	}
	return nil
}