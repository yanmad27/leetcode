package main

import (
	"encoding/xml"
	"fmt"
	"html"
	"reflect"
)

func expect(result, expectation interface{}) {
	if !reflect.DeepEqual(result, expectation) {
		panic(fmt.Sprintf("Expected %v, but got %v", expectation, result))
	}
}

type PatientParticipationListDocument struct {
	XMLName   xml.Name `xml:"PatientParticipationListDocument"`
	Text      string   `xml:",chardata"`
	Xsd       string   `xml:"xsd,attr"`
	Xsi       string   `xml:"xsi,attr"`
	ClassCode string   `xml:"classCode,attr"`
	MoodCode  string   `xml:"moodCode,attr"`
	Xmlns     string   `xml:"xmlns,attr"`
	RealmCode struct {
		Text string `xml:",chardata"`
		Code string `xml:"code,attr"`
	} `xml:"realmCode"`
	TypeId struct {
		Text      string `xml:",chardata"`
		Root      string `xml:"root,attr"`
		Extension string `xml:"extension,attr"`
	} `xml:"typeId"`
	TemplateId struct {
		Text string `xml:",chardata"`
		Root string `xml:"root,attr"`
	} `xml:"templateId"`
	ID struct {
		Text      string `xml:",chardata"`
		Root      string `xml:"root,attr"`
		Extension string `xml:"extension,attr"`
	} `xml:"id"`
	Code struct {
		Text        string `xml:",chardata"`
		Code        string `xml:"code,attr"`
		CodeSystem  string `xml:"codeSystem,attr"`
		DisplayName string `xml:"displayName,attr"`
	} `xml:"code"`
	Title         string `xml:"title"`
	EffectiveTime struct {
		Text string `xml:",chardata"`
		Low  struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"low"`
		High struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"high"`
	} `xml:"effectiveTime"`
	ConfidentialityCode struct {
		Text       string `xml:",chardata"`
		Code       string `xml:"code,attr"`
		CodeSystem string `xml:"codeSystem,attr"`
	} `xml:"confidentialityCode"`
	LanguageCode struct {
		Text string `xml:",chardata"`
		Code string `xml:"code,attr"`
	} `xml:"languageCode"`
	SetId struct {
		Text      string `xml:",chardata"`
		Root      string `xml:"root,attr"`
		Extension string `xml:"extension,attr"`
	} `xml:"setId"`
	VersionNumber struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"versionNumber"`
	Author struct {
		Text               string `xml:",chardata"`
		TypeCode           string `xml:"typeCode,attr"`
		ContextControlCode string `xml:"contextControlCode,attr"`
		Time               struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"time"`
		AssignedAuthor struct {
			Text      string `xml:",chardata"`
			ClassCode string `xml:"classCode,attr"`
			ID        struct {
				Text      string `xml:",chardata"`
				Root      string `xml:"root,attr"`
				Extension string `xml:"extension,attr"`
			} `xml:"id"`
			AssignedAuthoringDevice struct {
				Text                  string `xml:",chardata"`
				DeterminerCode        string `xml:"determinerCode,attr"`
				ManufacturerModelName string `xml:"manufacturerModelName"`
				SoftwareName          string `xml:"softwareName"`
			} `xml:"assignedAuthoringDevice"`
			RepresentedOrganization struct {
				Text           string `xml:",chardata"`
				ClassCode      string `xml:"classCode,attr"`
				DeterminerCode string `xml:"determinerCode,attr"`
				ID             struct {
					Text      string `xml:",chardata"`
					Root      string `xml:"root,attr"`
					Extension string `xml:"extension,attr"`
				} `xml:"id"`
				Name string `xml:"name"`
			} `xml:"representedOrganization"`
		} `xml:"assignedAuthor"`
	} `xml:"author"`
	Custodian struct {
		Text              string `xml:",chardata"`
		TypeCode          string `xml:"typeCode,attr"`
		AssignedCustodian struct {
			Text                             string `xml:",chardata"`
			ClassCode                        string `xml:"classCode,attr"`
			RepresentedCustodianOrganization struct {
				Text           string `xml:",chardata"`
				ClassCode      string `xml:"classCode,attr"`
				DeterminerCode string `xml:"determinerCode,attr"`
				ID             struct {
					Text string `xml:",chardata"`
					Root string `xml:"root,attr"`
				} `xml:"id"`
				Name string `xml:"name"`
			} `xml:"representedCustodianOrganization"`
		} `xml:"assignedCustodian"`
	} `xml:"custodian"`
	InformationRecipient struct {
		Text              string `xml:",chardata"`
		TypeCode          string `xml:"typeCode,attr"`
		IntendedRecipient struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text      string `xml:",chardata"`
				Root      string `xml:"root,attr"`
				Extension string `xml:"extension,attr"`
			} `xml:"id"`
			InformationRecipient struct {
				Text           string `xml:",chardata"`
				ClassCode      string `xml:"classCode,attr"`
				DeterminerCode string `xml:"determinerCode,attr"`
				Name           struct {
					Text   string `xml:",chardata"`
					Given  string `xml:"given"`
					Family string `xml:"family"`
					Prefix string `xml:"prefix"`
				} `xml:"name"`
			} `xml:"informationRecipient"`
		} `xml:"intendedRecipient"`
	} `xml:"informationRecipient"`
	Component struct {
		Text string `xml:",chardata"`
		Act  struct {
			Text       string `xml:",chardata"`
			ClassCode  string `xml:"classCode,attr"`
			MoodCode   string `xml:"moodCode,attr"`
			TemplateId struct {
				Text string `xml:",chardata"`
				Root string `xml:"root,attr"`
			} `xml:"templateId"`
			ID struct {
				Text      string `xml:",chardata"`
				Root      string `xml:"root,attr"`
				Extension string `xml:"extension,attr"`
			} `xml:"id"`
			Code struct {
				Text       string `xml:",chardata"`
				Code       string `xml:"code,attr"`
				CodeSystem string `xml:"codeSystem,attr"`
			} `xml:"code"`
			Title       string `xml:"title"`
			Participant []struct {
				Text               string `xml:",chardata"`
				TypeCode           string `xml:"typeCode,attr"`
				ContextControlCode string `xml:"contextControlCode,attr"`
				TemplateId         struct {
					Text string `xml:",chardata"`
					Root string `xml:"root,attr"`
				} `xml:"templateId"`
				Time struct {
					Text string `xml:",chardata"`
					Low  struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value,attr"`
					} `xml:"low"`
					High struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value,attr"`
					} `xml:"high"`
				} `xml:"time"`
				StatusCode struct {
					Text string `xml:",chardata"`
					Code string `xml:"code,attr"`
				} `xml:"statusCode"`
				ReasonCode struct {
					Text         string `xml:",chardata"`
					NullFlavor   string `xml:"nullFlavor,attr"`
					OriginalText string `xml:"originalText"`
				} `xml:"reasonCode"`
				AssociatedEntity struct {
					Text      string `xml:",chardata"`
					ClassCode string `xml:"classCode,attr"`
					ID        struct {
						Text       string `xml:",chardata"`
						Root       string `xml:"root,attr"`
						Extension  string `xml:"extension,attr"`
						NullFlavor string `xml:"nullFlavor,attr"`
					} `xml:"id"`
					AssociatedPatient struct {
						Text           string `xml:",chardata"`
						ClassCode      string `xml:"classCode,attr"`
						DeterminerCode string `xml:"determinerCode,attr"`
						Name           struct {
							Text   string `xml:",chardata"`
							Family struct {
								Text       string `xml:",chardata"`
								NullFlavor string `xml:"nullFlavor,attr"`
							} `xml:"family"`
							Given struct {
								Text       string `xml:",chardata"`
								NullFlavor string `xml:"nullFlavor,attr"`
							} `xml:"given"`
						} `xml:"name"`
						AdministrativeGenderCode struct {
							Text       string `xml:",chardata"`
							Code       string `xml:"code,attr"`
							CodeSystem string `xml:"codeSystem,attr"`
						} `xml:"administrativeGenderCode"`
						BirthTime struct {
							Text       string `xml:",chardata"`
							Value      string `xml:"value,attr"`
							NullFlavor string `xml:"nullFlavor,attr"`
						} `xml:"birthTime"`
					} `xml:"associatedPatient"`
					ScopingOrganization struct {
						Text           string `xml:",chardata"`
						ClassCode      string `xml:"classCode,attr"`
						DeterminerCode string `xml:"determinerCode,attr"`
						ID             struct {
							Text      string `xml:",chardata"`
							Root      string `xml:"root,attr"`
							Extension string `xml:"extension,attr"`
						} `xml:"id"`
						Name struct {
							Text       string `xml:",chardata"`
							NullFlavor string `xml:"nullFlavor,attr"`
						} `xml:"name"`
					} `xml:"scopingOrganization"`
				} `xml:"associatedEntity"`
			} `xml:"participant"`
		} `xml:"act"`
	} `xml:"component"`
}

func main() {
	str := "&lt;?xml version=&quot;1.0&quot; encoding=&quot;utf-8&quot;?&gt;\r\n&lt;PatientParticipationListDocument xmlns:xsd=&quot;http://www.w3.org/2001/XMLSchema&quot; xmlns:xsi=&quot;http://www.w3.org/2001/XMLSchema-instance&quot; classCode=&quot;DOC&quot; moodCode=&quot;EVN&quot; xmlns=&quot;urn:hl7-org:v3&quot;&gt;\r\n  &lt;realmCode code=&quot;DE&quot; /&gt;\r\n  &lt;typeId root=&quot;2.16.840.1.113883.1.3&quot; extension=&quot;POMF_HD000001DE&quot; /&gt;\r\n  &lt;templateId root=&quot;1.2.276.0.76.10.1018&quot; /&gt;\r\n  &lt;id root=&quot;1.2.276.0.76.3.1.217.1876766&quot; extension=&quot;159df055-0157-4d65-8558-067dd8da747b&quot; /&gt;\r\n  &lt;code code=&quot;64291-8&quot; codeSystem=&quot;2.16.840.1.113883.6.1&quot; displayName=&quot;Health insurance-related form&quot; /&gt;\r\n  &lt;title&gt;Patiententeilnehmerliste&lt;/title&gt;\r\n  &lt;effectiveTime&gt;\r\n    &lt;low value=&quot;20250701000000&quot; /&gt;\r\n    &lt;high value=&quot;20250930235959&quot; /&gt;\r\n  &lt;/effectiveTime&gt;\r\n  &lt;confidentialityCode code=&quot;N&quot; codeSystem=&quot;2.16.840.1.113883.5.25&quot; /&gt;\r\n  &lt;languageCode code=&quot;de-DE&quot; /&gt;\r\n  &lt;setId root=&quot;1.2.276.0.76.3.1.217.1876767&quot; extension=&quot;370124e3-e5bb-4f21-8ec5-b708fb8a6121&quot; /&gt;\r\n  &lt;versionNumber value=&quot;1&quot; /&gt;\r\n  &lt;author typeCode=&quot;AUT&quot; contextControlCode=&quot;OP&quot;&gt;\r\n    &lt;time value=&quot;20170915125232&quot; /&gt;\r\n    &lt;assignedAuthor classCode=&quot;ASSIGNED&quot;&gt;\r\n      &lt;id root=&quot;1.2.276.0.76.4.5&quot; extension=&quot;304208930&quot; /&gt;\r\n      &lt;assignedAuthoringDevice determinerCode=&quot;INSTANCE&quot;&gt;\r\n        &lt;manufacturerModelName&gt;HAEVG_Patiententeilnehmerliste&lt;/manufacturerModelName&gt;\r\n        &lt;softwareName&gt;HAEVG_Patiententeilnehmerliste&lt;/softwareName&gt;\r\n      &lt;/assignedAuthoringDevice&gt;\r\n      &lt;representedOrganization classCode=&quot;ORG&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n        &lt;id root=&quot;1.2.276.0.76.4.5&quot; extension=&quot;304208930&quot; /&gt;\r\n        &lt;name&gt;H&#196;VG Rechenzentrum GmbH&lt;/name&gt;\r\n      &lt;/representedOrganization&gt;\r\n    &lt;/assignedAuthor&gt;\r\n  &lt;/author&gt;\r\n  &lt;custodian typeCode=&quot;CST&quot;&gt;\r\n    &lt;assignedCustodian classCode=&quot;ASSIGNED&quot;&gt;\r\n      &lt;representedCustodianOrganization classCode=&quot;ORG&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n        &lt;id root=&quot;1.2.276.0.76.3.1.217&quot; /&gt;\r\n        &lt;name&gt;H&#196;VG AG&lt;/name&gt;\r\n      &lt;/representedCustodianOrganization&gt;\r\n    &lt;/assignedCustodian&gt;\r\n  &lt;/custodian&gt;\r\n  &lt;informationRecipient typeCode=&quot;PRCP&quot;&gt;\r\n    &lt;intendedRecipient&gt;\r\n      &lt;id root=&quot;1.2.276.0.76.4.16&quot; extension=&quot;999956201&quot; /&gt;\r\n      &lt;id root=&quot;1.2.276.0.76.3.1.217.4.2&quot; extension=&quot;H4AAAFK&quot; /&gt;\r\n      &lt;informationRecipient classCode=&quot;PSN&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n        &lt;name&gt;\r\n          &lt;given&gt;Karoline&lt;/given&gt;\r\n          &lt;family&gt;Arzt K&lt;/family&gt;\r\n          &lt;prefix&gt;Dr. med.&lt;/prefix&gt;\r\n        &lt;/name&gt;\r\n      &lt;/informationRecipient&gt;\r\n      &lt;receivedOrganization classCode=&quot;ORG&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n        &lt;name&gt;Vertragspartner H4AAAFK&lt;/name&gt;\r\n      &lt;/receivedOrganization&gt;\r\n    &lt;/intendedRecipient&gt;\r\n  &lt;/informationRecipient&gt;\r\n  &lt;component&gt;\r\n    &lt;act classCode=&quot;CNTRCT&quot; moodCode=&quot;EVN&quot;&gt;\r\n      &lt;templateId root=&quot;1.2.276.0.76.10.4081&quot; /&gt;\r\n      &lt;id root=&quot;1.2.276.0.76.3.1.217.4.1&quot; extension=&quot;AOK_HE_HZV&quot; /&gt;\r\n      &lt;code code=&quot;64291-8&quot; codeSystem=&quot;2.16.840.1.113883.6.1&quot; /&gt;\r\n      &lt;title&gt;AOK Hessen hausarztzentrierte Versorgung&lt;/title&gt;\r\n      &lt;participant typeCode=&quot;COV&quot; contextControlCode=&quot;OP&quot;&gt;\r\n        &lt;templateId root=&quot;1.2.276.0.76.10.4082&quot; /&gt;\r\n        &lt;time&gt;\r\n          &lt;low value=&quot;20250701000000&quot; /&gt;\r\n        &lt;/time&gt;\r\n        &lt;statusCode code=&quot;active&quot; /&gt;\r\n        &lt;associatedEntity classCode=&quot;PAT&quot;&gt;\r\n          &lt;id root=&quot;1.2.276.0.76.4.8&quot; extension=&quot;D441503770&quot; /&gt;\r\n          &lt;associatedPatient classCode=&quot;PSN&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n            &lt;name&gt;\r\n              &lt;family&gt;FISCHER&lt;/family&gt;\r\n              &lt;given&gt;PHILIPP&lt;/given&gt;\r\n            &lt;/name&gt;\r\n            &lt;administrativeGenderCode code=&quot;M&quot; codeSystem=&quot;2.16.840.1.113883.5.1&quot; /&gt;\r\n            &lt;birthTime value=&quot;20100729&quot; /&gt;\r\n          &lt;/associatedPatient&gt;\r\n          &lt;scopingOrganization classCode=&quot;ORG&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n            &lt;id root=&quot;1.2.276.0.76.4.5&quot; extension=&quot;105513158&quot; /&gt;\r\n            &lt;name&gt;Krankenkasse IK:105513158&lt;/name&gt;\r\n          &lt;/scopingOrganization&gt;\r\n        &lt;/associatedEntity&gt;\r\n      &lt;/participant&gt;\r\n      &lt;participant typeCode=&quot;COV&quot; contextControlCode=&quot;OP&quot;&gt;\r\n        &lt;templateId root=&quot;1.2.276.0.76.10.4082&quot; /&gt;\r\n        &lt;time&gt;\r\n          &lt;low value=&quot;20250701000000&quot; /&gt;\r\n        &lt;/time&gt;\r\n        &lt;statusCode code=&quot;rejected&quot; /&gt;\r\n        &lt;reasonCode nullFlavor=&quot;OTH&quot;&gt;\r\n          &lt;originalText&gt;Kein Versicherungsverh�ltnis bekannt&lt;/originalText&gt;\r\n        &lt;/reasonCode&gt;\r\n        &lt;associatedEntity classCode=&quot;PAT&quot;&gt;\r\n          &lt;id root=&quot;1.2.276.0.76.4.8&quot; extension=&quot;M499311275&quot; /&gt;\r\n          &lt;associatedPatient classCode=&quot;PSN&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n            &lt;name&gt;\r\n              &lt;family&gt;WERNER&lt;/family&gt;\r\n              &lt;given&gt;JONAS&lt;/given&gt;\r\n            &lt;/name&gt;\r\n            &lt;administrativeGenderCode code=&quot;M&quot; codeSystem=&quot;2.16.840.1.113883.5.1&quot; /&gt;\r\n            &lt;birthTime value=&quot;20111125&quot; /&gt;\r\n          &lt;/associatedPatient&gt;\r\n          &lt;scopingOrganization classCode=&quot;ORG&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n            &lt;id root=&quot;1.2.276.0.76.4.5&quot; extension=&quot;105513136&quot; /&gt;\r\n            &lt;name&gt;Krankenkasse IK:105513136&lt;/name&gt;\r\n          &lt;/scopingOrganization&gt;\r\n        &lt;/associatedEntity&gt;\r\n      &lt;/participant&gt;\r\n      &lt;participant typeCode=&quot;COV&quot; contextControlCode=&quot;OP&quot;&gt;\r\n        &lt;templateId root=&quot;1.2.276.0.76.10.4082&quot; /&gt;\r\n        &lt;time&gt;\r\n          &lt;low value=&quot;20250701000000&quot; /&gt;\r\n        &lt;/time&gt;\r\n        &lt;statusCode code=&quot;requested&quot; /&gt;\r\n        &lt;associatedEntity classCode=&quot;PAT&quot;&gt;\r\n          &lt;id root=&quot;1.2.276.0.76.4.8&quot; extension=&quot;H617866491&quot; /&gt;\r\n          &lt;associatedPatient classCode=&quot;PSN&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n            &lt;name&gt;\r\n              &lt;family&gt;LEHMANN&lt;/family&gt;\r\n              &lt;given&gt;MIA&lt;/given&gt;\r\n            &lt;/name&gt;\r\n            &lt;administrativeGenderCode code=&quot;F&quot; codeSystem=&quot;2.16.840.1.113883.5.1&quot; /&gt;\r\n            &lt;birthTime value=&quot;19351101&quot; /&gt;\r\n          &lt;/associatedPatient&gt;\r\n          &lt;scopingOrganization classCode=&quot;ORG&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n            &lt;id root=&quot;1.2.276.0.76.4.5&quot; extension=&quot;105413179&quot; /&gt;\r\n            &lt;name&gt;Krankenkasse IK:105413179&lt;/name&gt;\r\n          &lt;/scopingOrganization&gt;\r\n        &lt;/associatedEntity&gt;\r\n      &lt;/participant&gt;\r\n      &lt;participant typeCode=&quot;COV&quot; contextControlCode=&quot;OP&quot;&gt;\r\n        &lt;templateId root=&quot;1.2.276.0.76.10.4082&quot; /&gt;\r\n        &lt;time&gt;\r\n          &lt;low value=&quot;20230701000000&quot; /&gt;\r\n          &lt;high value=&quot;20250930235959&quot; /&gt;\r\n        &lt;/time&gt;\r\n        &lt;statusCode code=&quot;active&quot; /&gt;\r\n        &lt;reasonCode nullFlavor=&quot;OTH&quot;&gt;\r\n          &lt;originalText&gt;K�ndigung von HzV&lt;/originalText&gt;\r\n        &lt;/reasonCode&gt;\r\n        &lt;associatedEntity classCode=&quot;PAT&quot;&gt;\r\n          &lt;id root=&quot;1.2.276.0.76.4.8&quot; extension=&quot;J698131734&quot; /&gt;\r\n          &lt;associatedPatient classCode=&quot;PSN&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n            &lt;name&gt;\r\n              &lt;family&gt;ALBRECHT&lt;/family&gt;\r\n              &lt;given&gt;MAXIMILIAN&lt;/given&gt;\r\n            &lt;/name&gt;\r\n            &lt;administrativeGenderCode code=&quot;M&quot; codeSystem=&quot;2.16.840.1.113883.5.1&quot; /&gt;\r\n            &lt;birthTime value=&quot;19750315&quot; /&gt;\r\n          &lt;/associatedPatient&gt;\r\n          &lt;scopingOrganization classCode=&quot;ORG&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n            &lt;id root=&quot;1.2.276.0.76.4.5&quot; extension=&quot;105713207&quot; /&gt;\r\n            &lt;name&gt;Krankenkasse IK:105713207&lt;/name&gt;\r\n          &lt;/scopingOrganization&gt;\r\n        &lt;/associatedEntity&gt;\r\n      &lt;/participant&gt;\r\n      &lt;participant typeCode=&quot;COV&quot; contextControlCode=&quot;OP&quot;&gt;\r\n        &lt;templateId root=&quot;1.2.276.0.76.10.4082&quot; /&gt;\r\n        &lt;time&gt;\r\n          &lt;low value=&quot;20250401000000&quot; /&gt;\r\n        &lt;/time&gt;\r\n        &lt;statusCode code=&quot;active&quot; /&gt;\r\n        &lt;associatedEntity classCode=&quot;PAT&quot;&gt;\r\n          &lt;id root=&quot;1.2.276.0.76.4.8&quot; extension=&quot;B637895240&quot; /&gt;\r\n          &lt;associatedPatient classCode=&quot;PSN&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n            &lt;name&gt;\r\n              &lt;family&gt;GRUENHEID&lt;/family&gt;\r\n              &lt;given&gt;KARL&lt;/given&gt;\r\n            &lt;/name&gt;\r\n            &lt;administrativeGenderCode code=&quot;M&quot; codeSystem=&quot;2.16.840.1.113883.5.1&quot; /&gt;\r\n            &lt;birthTime value=&quot;19400312&quot; /&gt;\r\n          &lt;/associatedPatient&gt;\r\n          &lt;scopingOrganization classCode=&quot;ORG&quot; determinerCode=&quot;INSTANCE&quot;&gt;\r\n            &lt;id root=&quot;1.2.276.0.76.4.5&quot; extension=&quot;105513136&quot; /&gt;\r\n            &lt;name&gt;Krankenkasse IK:105513136&lt;/name&gt;\r\n          &lt;/scopingOrganization&gt;\r\n        &lt;/associatedEntity&gt;\r\n      &lt;/participant&gt;\r\n    &lt;/act&gt;\r\n  &lt;/component&gt;\r\n&lt;/PatientParticipationListDocument&gt;"
	unescaped := html.UnescapeString(str)
	fmt.Println(unescaped)

	patientPartipationList := PatientParticipationListDocument{}
	if err := xml.Unmarshal([]byte(unescaped), &patientPartipationList); err != nil {
		panic("ee")
	}
	expect(unescaped, "Kündigung von HZV")
}
