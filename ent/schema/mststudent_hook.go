package schema

import (
	"adm-num/ent"
	"adm-num/ent/hook"
	"context"
	"fmt"
	"strconv"
)

func MstStudentGetAdmissionNumber(next ent.Mutator) ent.Mutator {
	return hook.MstStudentFunc(func(ctx context.Context, m *ent.MstStudentMutation) (ent.Value, error) {

		client := m.Client()

		var AdmNumber string

		AdmNumber, err := GetInstAdmNextNumber(ctx, client)
		if err != nil {
			return nil, err
		}

		m.SetStdAdmNo(AdmNumber)

		v, verr := next.Mutate(ctx, m)

		if verr != nil {

			return v, verr
		}

		admNumRec, err := client.AdmNumber.Query().Only(ctx)
		if err != nil {

			return v, err
		}

		admNum, err := strconv.Atoi(admNumRec.AdmCurrentNo)

		if err != nil {

			return v, err
		}

		nextAdmNum := fmt.Sprintf("%d", admNum+1)

		/* update the voucher number */
		_, err = client.AdmNumber.UpdateOneID(admNumRec.ID).
			SetAdmCurrentNo(nextAdmNum).
			Save(ctx)

		if err != nil {

			return v, err
		}

		return v, verr

	})
}

func GetInstAdmNextNumber(ctx context.Context, client *ent.Client) (string, error) {

	instAdmNumberRecord, err := client.AdmNumber.Query().Only(ctx)

	if err != nil {
		return "", err
	}

	admNum := instAdmNumberRecord.AdmCurrentNo

	if instAdmNumberRecord.PrefillWithZero && instAdmNumberRecord.PrefillWidth > 0 {
		num, err := strconv.Atoi(admNum)
		if err != nil {
			return "", err
		}
		switch instAdmNumberRecord.PrefillWidth {
		case 1:
			admNum = fmt.Sprintf("%01d", num+1)

		case 2:
			admNum = fmt.Sprintf("%02d", num+1)

		case 3:
			admNum = fmt.Sprintf("%03d", num+1)

		case 4:
			admNum = fmt.Sprintf("%04d", num+1)

		case 5:
			admNum = fmt.Sprintf("%05d", num+1)

		}
	} else {
		num, _ := strconv.Atoi(admNum)
		admNum = fmt.Sprintf("%d", num+1)
	}
	var finalAdmNumber string
	if instAdmNumberRecord.IsPrefixed {
		if len(instAdmNumberRecord.PrefixStr) > 0 {
			finalAdmNumber = instAdmNumberRecord.PrefixStr
		}
		if len(instAdmNumberRecord.Separator) > 0 {
			finalAdmNumber = finalAdmNumber + instAdmNumberRecord.Separator
		}
		finalAdmNumber = finalAdmNumber + admNum

		if len(instAdmNumberRecord.SuffixStr) > 0 {
			if len(instAdmNumberRecord.Separator) > 0 {
				finalAdmNumber = finalAdmNumber + instAdmNumberRecord.Separator
			}
			finalAdmNumber = finalAdmNumber + instAdmNumberRecord.SuffixStr
		}

	} else {
		finalAdmNumber = admNum
	}
	return finalAdmNumber, nil
}
