package delivery

import (
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "go/pkg/services/contact/internal/usecase"
    "log"
    "net/http"
)

// ContactHandler represents the HTTP handler for contact operations
type ContactHandler struct {
    useCase usecase.ContactUseCase
    logger  *log.Logger
}

// NewContactHandler creates a new instance of ContactHandler
func NewContactHandler(useCase usecase.ContactUseCase, logger *log.Logger) *ContactHandler {
    return &ContactHandler{
        useCase: useCase,
        logger:  logger,
    }
}

// HandleHTTP handles HTTP requests for contact operations
func (h *ContactHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
    ctx := context.WithValue(r.Context(), "requestID", uuid.New())

    // Log request information
    h.logger.Printf("%s %s %s\n", r.Method, r.URL.Path, r.Proto)

    switch r.Method {
    case http.MethodGet:
        h.logger.Println("GET /contact")
        h.getContact(w, r.WithContext(ctx))
    case http.MethodPost:
        h.logger.Println("POST /contact")
        h.createContact(w, r.WithContext(ctx))
    case http.MethodPut:
        h.logger.Println("PUT /contact")
        h.updateContact(w, r.WithContext(ctx))
    case http.MethodDelete:
        h.logger.Println("DELETE /contact")
        h.deleteContact(w, r.WithContext(ctx))
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// getContact handles HTTP GET requests to retrieve a contact
func (h *ContactHandler) getContact(w http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value("requestID").(uuid.UUID)

	contact, err := h.useCase.GetContact(requestID)
	if err != nil {
		log.Printf("Error getting contact: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}

// createContact handles HTTP POST requests to create a new contact
func (h *ContactHandler) createContact(w http.ResponseWriter, r *http.Request) {
    var contact usecase.Contact
    err := json.NewDecoder(r.Body).Decode(&contact)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	err = h.useCase.CreateContact(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}

// updateContact handles HTTP PUT requests to update an existing contact
func (h *ContactHandler) updateContact(w http.ResponseWriter, r *http.Request) {
    var contact usecase.Contact
    err := json.NewDecoder(r.Body).Decode(&contact)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = h.useCase.UpdateContact(r.Context(), &contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}

// deleteContact handles HTTP DELETE requests to delete a contact
func (h *ContactHandler) deleteContact(w http.ResponseWriter, r *http.Request) {
    err := h.useCase.DeleteContact(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}

type GroupHandler struct {
	useCase usecase.GroupUseCase
}

func NewGroupHandler(useCase usecase.GroupUseCase) *GroupHandler {
	return &GroupHandler{
		useCase: useCase,
	}
}

func (h *GroupHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "requestID", uuid.New())

	switch r.Method {
	case http.MethodGet:
		h.getGroup(w, r.WithContext(ctx))
	case http.MethodPost:
		h.createGroup(w, r.WithContext(ctx))
	case http.MethodPut:
		h.updateGroup(w, r.WithContext(ctx))
	case http.MethodDelete:
		h.deleteGroup(w, r.WithContext(ctx))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *GroupHandler) getGroup(w http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value("requestID").(uuid.UUID)

	group, err := h.useCase.GetGroup(requestID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) createGroup(w http.ResponseWriter, r *http.Request) {
	var group usecase.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.useCase.CreateGroup(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) updateGroup(w http.ResponseWriter, r *http.Request) {
	var group usecase.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.useCase.UpdateGroup(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) deleteGroup(w http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value("requestID").(uuid.UUID)

	err := h.useCase.DeleteGroup(requestID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}
