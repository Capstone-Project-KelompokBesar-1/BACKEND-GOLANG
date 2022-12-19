package helpers

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func InitMidtransSnap() snap.Client {
	var s snap.Client
	s.New("SB-Mid-server-Fq8L95I4SLW5dwYtw6FsuTud", midtrans.Sandbox)

	return s
}
