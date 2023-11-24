package repository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"taps/domain"
	"taps/domain/vo"
	"taps/utils/clk"
	"testing"
	"time"
)

func TestInfluenzaCensusStore(t *testing.T) {
	t.Run("Save", func(t *testing.T) {
		t.Run("saves a census", func(t *testing.T) {
			dynamoDB, table := setupInfluenzaCensusStore()
			defer dynamoDB.DeleteTable(&dynamodb.DeleteTableInput{TableName: aws.String(table)})
			influenzaCensusStore := NewDynamoInfluenzaCensusRepository(
				dynamoDB,
				table,
				clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
			)

			err := influenzaCensusStore.Save(domain.Census{
				ID: "RAHE190116MMCMRSA7",
				CURP: vo.Curp{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					State:         "MEXICO",
					Number:        15,
				},
				Address:         AddressFixture(t, "18b", "chapulin", "arcoiris"),
				ApplicationDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				TargetGroup:     vo.MustNewTargetGroup(true, false),
				RiskGroup: vo.RiskGroup{
					PregnantWomen:                           true,
					WellnessPerson:                          true,
					AIDS:                                    true,
					Diabetes:                                true,
					Obesity:                                 true,
					AcuteOrChronicHeartDisease:              true,
					ChronicLungDiseaseIncludesCOPDAndAsthma: true,
					Cancer:                                  true,
					ChronicConditionsThatRequireProlongedConsumptionOfSalicylic: true,
					RenalInsufficiency: true,
					AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS: true,
					EssentialHypertension: true,
				},
				SeasonalInfluenzaVaccinationSchedule: vo.MustNewSeasonalInfluenzaVaccinationSchedule(true, false, false),
				Rights:                               vo.Rights.ISSSTE,
			})

			require.NoError(t, err)
		})

		t.Run("returns ERROR when providing census with empty ID", func(t *testing.T) {
			dynamoDB, table := setupInfluenzaCensusStore()
			defer dynamoDB.DeleteTable(&dynamodb.DeleteTableInput{TableName: aws.String(table)})
			influenzaCensusStore := NewDynamoInfluenzaCensusRepository(
				dynamoDB,
				table,
				clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
			)

			err := influenzaCensusStore.Save(domain.Census{
				ID: "",
				CURP: vo.Curp{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					State:         "MEXICO",
					Number:        15,
				},
				Address:         AddressFixture(t, "18b", "chapulin", "arcoiris"),
				ApplicationDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				TargetGroup:     vo.MustNewTargetGroup(true, false),
				RiskGroup: vo.RiskGroup{
					PregnantWomen:                           true,
					WellnessPerson:                          true,
					AIDS:                                    true,
					Diabetes:                                true,
					Obesity:                                 true,
					AcuteOrChronicHeartDisease:              true,
					ChronicLungDiseaseIncludesCOPDAndAsthma: true,
					Cancer:                                  true,
					ChronicConditionsThatRequireProlongedConsumptionOfSalicylic: true,
					RenalInsufficiency: true,
					AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS: true,
					EssentialHypertension: true,
				},
				SeasonalInfluenzaVaccinationSchedule: vo.MustNewSeasonalInfluenzaVaccinationSchedule(true, false, false),
				Rights:                               vo.Rights.ISSSTE,
			})

			require.ErrorContains(t, err, "ValidationException: Invalid attribute value type")
		})
	})

	t.Run("Find", func(t *testing.T) {
		t.Run("returns TRUE when the census is found", func(t *testing.T) {
			dynamoDB, table := setupInfluenzaCensusStore()
			defer dynamoDB.DeleteTable(&dynamodb.DeleteTableInput{TableName: aws.String(table)})
			influenzaCensusStore := NewDynamoInfluenzaCensusRepository(
				dynamoDB,
				table,
				clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
			)

			err := influenzaCensusStore.Save(
				domain.Census{
					ID: "RAHE190116MMCMRSA7",
					CURP: vo.Curp{
						ID:            "RAHE190116MMCMRSA7",
						LastLastName:  "RAMIREZ",
						FirstLastName: "HERRERA",
						FirstName:     "ESTHER ELIZABETH",
						Gender:        "MUJER",
						DOB:           "16/01/2019",
						State:         "MEXICO",
						Number:        15,
					},
					Address:         AddressFixture(t, "18b", "chapulin", "arcoiris"),
					ApplicationDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					TargetGroup:     vo.MustNewTargetGroup(true, false),
					RiskGroup: vo.RiskGroup{
						PregnantWomen:                           true,
						WellnessPerson:                          true,
						AIDS:                                    true,
						Diabetes:                                true,
						Obesity:                                 true,
						AcuteOrChronicHeartDisease:              true,
						ChronicLungDiseaseIncludesCOPDAndAsthma: true,
						Cancer:                                  true,
						ChronicConditionsThatRequireProlongedConsumptionOfSalicylic: true,
						RenalInsufficiency: true,
						AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS: true,
						EssentialHypertension: true,
					},
					SeasonalInfluenzaVaccinationSchedule: vo.MustNewSeasonalInfluenzaVaccinationSchedule(true, false, false),
					Rights:                               vo.Rights.ISSSTE,
				},
			)
			census, found, err := influenzaCensusStore.Find("RAHE190116MMCMRSA7", "2023-01-01")

			require.NoError(t, err)
			require.True(t, found)
			require.Equal(
				t,
				census,
				domain.Census{
					ID: "RAHE190116MMCMRSA7",
					CURP: vo.Curp{
						ID:            "RAHE190116MMCMRSA7",
						LastLastName:  "RAMIREZ",
						FirstLastName: "HERRERA",
						FirstName:     "ESTHER ELIZABETH",
						Gender:        "MUJER",
						DOB:           "16/01/2019",
						State:         "MEXICO",
						Number:        15,
					},
					Address:         AddressFixture(t, "18b", "chapulin", "arcoiris"),
					ApplicationDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					TargetGroup:     vo.MustNewTargetGroup(true, false),
					RiskGroup: vo.RiskGroup{
						PregnantWomen:                           true,
						WellnessPerson:                          true,
						AIDS:                                    true,
						Diabetes:                                true,
						Obesity:                                 true,
						AcuteOrChronicHeartDisease:              true,
						ChronicLungDiseaseIncludesCOPDAndAsthma: true,
						Cancer:                                  true,
						ChronicConditionsThatRequireProlongedConsumptionOfSalicylic: true,
						RenalInsufficiency: true,
						AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS: true,
						EssentialHypertension: true,
					},
					SeasonalInfluenzaVaccinationSchedule: vo.MustNewSeasonalInfluenzaVaccinationSchedule(true, false, false),
					Rights:                               vo.Rights.ISSSTE,
				},
			)
		})

		t.Run("returns FALSE when the census is not found", func(t *testing.T) {
			dynamoDB, table := setupInfluenzaCensusStore()
			defer dynamoDB.DeleteTable(&dynamodb.DeleteTableInput{TableName: aws.String(table)})

			influenzaCensusStore := NewDynamoInfluenzaCensusRepository(dynamoDB, table, clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)))
			census, found, err := influenzaCensusStore.Find("NOT-EXISTING-CENSUS", "2023-1-1")

			require.NoError(t, err)
			require.False(t, found)
			require.Equal(t, census, domain.Census{})
		})

		t.Run("returns ERROR when the repository returns an error", func(t *testing.T) {
			dynamoDB, table := setupInfluenzaCensusStore()
			defer dynamoDB.DeleteTable(&dynamodb.DeleteTableInput{TableName: aws.String(table)})
			influenzaCensusStore := NewDynamoInfluenzaCensusRepository(dynamoDB, table, clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)))
			census, found, err := influenzaCensusStore.Find("", "")

			require.Equal(t, found, false)
			require.ErrorContains(
				t,
				err,
				"ValidationException: One or more parameter values are not valid. The AttributeValue for a key attribute cannot contain an empty string value.",
			)
			require.Equal(t, census, domain.Census{})
		})
	})
}

func setupInfluenzaCensusStore() (*dynamodb.DynamoDB, string) {
	dynamoDB := dynamodb.New(
		session.Must(
			session.NewSession(
				&aws.Config{
					Region:      aws.String("local"),
					Credentials: credentials.NewStaticCredentials("id", "secret", "token"),
					Endpoint:    aws.String("http://localhost:8000"),
				},
			),
		),
	)

	table := fmt.Sprintf("influenza-census-%s", uuid.New().String())

	_, err := dynamoDB.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(table),
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("CurpID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Date"),
				KeyType:       aws.String("RANGE"),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("CurpID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Date"),
				AttributeType: aws.String("S"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	})

	if err != nil {
		panic(err)
	}

	return dynamoDB, table
}

func AddressFixture(t *testing.T, streetNumber, streetName, suburbName string) vo.Address {
	address, err := vo.TryNewAddress(streetNumber, streetName, suburbName)
	require.NoError(t, err)
	return address
}
