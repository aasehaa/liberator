package nais_io_v1alpha1

import (
	"time"

	"github.com/nais/liberator/pkg/stringutil"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const eventMaxLen = 1024

func (in *Application) CreateEvent(reason, message, typeStr string) *corev1.Event {
	objectMeta := in.CreateObjectMeta()
	objectMeta.GenerateName = "naiserator-event-"
	objectMeta.Name = ""

	return &corev1.Event{
		ObjectMeta:          objectMeta,
		ReportingController: "naiserator",
		ReportingInstance:   "naiserator",
		Action:              reason,
		Reason:              reason,
		InvolvedObject:      in.GetObjectReference(),
		Source:              corev1.EventSource{Component: "naiserator"},
		Message:             stringutil.StrTrimMiddle(message, eventMaxLen),
		EventTime:           metav1.MicroTime{Time: time.Now()},
		FirstTimestamp:      metav1.Time{Time: time.Now()},
		LastTimestamp:       metav1.Time{Time: time.Now()},
		Type:                typeStr,
		Count:               1,
	}
}
