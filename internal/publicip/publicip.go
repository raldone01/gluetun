package publicip

import (
	"github.com/qdm12/gluetun/internal/publicip/models"
)

func (l *Loop) GetData() (data models.IPInfoData) {
	return l.state.GetData()
}
