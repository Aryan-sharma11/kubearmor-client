package vm

import (
	"fmt"

	pb "github.com/kubearmor/KubeArmor/protobuf"
)

func responseHandler(resp *pb.Response, eventType string) error {

	switch resp.Status {

	case 0, 4, -3:
		return fmt.Errorf(" Failed ")

	case 1:
		fmt.Println("Success")

	case 2:
		fmt.Println(" Policy applied successfully ")

	case 3:
		fmt.Println(" Policy configured ")

	case -1:
		fmt.Println(" Policy deleted ")

	case -2:
		return fmt.Errorf(" Policy doesn't exist ")

	}
	return nil

}
