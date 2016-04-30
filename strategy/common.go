package strategy

import (
	"github.com/xh3b4sd/anna/smart-map"
	"github.com/xh3b4sd/anna/spec"
)

func stringMapToStrategy(stringMap map[string]string) (*strategy, error) {
	newConfig := smartmap.DefaultConfig()
	newConfig.StringMap = stringMap
	newSmartMap, err := smartmap.NewSmartMap(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	newStrategy := &strategy{}
	newStrategy.CLGNames, err = newSmartMap.GetStringStringSlice("clg-names")
	if err != nil {
		return nil, maskAny(err)
	}
	newID, err := newSmartMap.GetStringString("id")
	if err != nil {
		return nil, maskAny(err)
	}
	newStrategy.ID = spec.ObjectID(newID)
	newRequestor, err := newSmartMap.GetStringString("requestor")
	if err != nil {
		return nil, maskAny(err)
	}
	newStrategy.Requestor = spec.ObjectType(newRequestor)
	newType, err := newSmartMap.GetStringString("type")
	if err != nil {
		return nil, maskAny(err)
	}
	newStrategy.Type = spec.ObjectType(newType)

	return newStrategy, nil
}

func strategyToStringMap(strat *strategy) (map[string]string, error) {
	newSmartMap, err := smartmap.NewSmartMap(smartmap.DefaultConfig())
	if err != nil {
		return nil, maskAny(err)
	}

	err = newSmartMap.SetStringStringSlice("clg-names", strat.GetCLGNames())
	if err != nil {
		return nil, maskAny(err)
	}
	err = newSmartMap.SetStringString("id", string(strat.GetID()))
	if err != nil {
		return nil, maskAny(err)
	}
	err = newSmartMap.SetStringString("requestor", string(strat.GetRequestor()))
	if err != nil {
		return nil, maskAny(err)
	}
	err = newSmartMap.SetStringString("type", string(strat.GetType()))
	if err != nil {
		return nil, maskAny(err)
	}

	return newSmartMap.GetStringMap(), nil
}
