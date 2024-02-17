package delivery

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"go/pkg/services/contact/internal/usecase"
	"log"
	"net/http"
)


type ContactHandler struct {
	useCase usecase.ContactUseCase
	logger  *log.Logger
}


func NewContactHandler(useCase usecase.ContactUseCase, logger *log.Logger) *ContactHandler {
	return &ContactHandler{
		useCase: useCase,
		logger:  logger,
	}
}


func (h *ContactHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
    
    requestID := uuid.New()

    
    traceID := r.Header.Get("X-Trace-ID")
    if traceID == "" {
        traceID = uuid.New().String()
    }

    
    ctx := context.WithValue(r.Context(), "traceID", traceID)
    ctx = context.WithValue(ctx, "requestID", requestID)

	
	h.logger.Printf("[%s] %s %s %s\n", traceID, r.Method, r.URL.Path, r.Proto)

	
	switch r.Method {
	case http.MethodGet:
		h.getContact(w, r.WithContext(ctx))
	case http.MethodPost:
		h.createContact(w, r.WithContext(ctx))
	case http.MethodPut:
		h.updateContact(w, r.WithContext(ctx))
	case http.MethodDelete:
		h.deleteContact(w, r.WithContext(ctx))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}



func (h *ContactHandler) getContact(w http.ResponseWriter, r *http.Request) {
	
	traceID := r.Context().Value("traceID").(string)

	
	h.logger.Printf("[%s] Getting contact\n", traceID)

	
	requestID := r.Context().Value("requestID").(uuid.UUID)

	
	contact, err := h.useCase.GetContact(requestID)
	if err != nil {
		
		h.logger.Printf("[%s] Error getting contact: %v\n", traceID, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}


func (h *ContactHandler) createContact(w http.ResponseWriter, r *http.Request) {
	
	traceID := r.Context().Value("traceID").(string)

	
	h.logger.Printf("[%s] Creating contact\n", traceID)

	
	var contact usecase.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		
		h.logger.Printf("[%s] Error decoding request body: %v\n", traceID, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	
	err = h.useCase.CreateContact(&contact)
	if err != nil {
		
		h.logger.Printf("[%s] Error creating contact: %v\n", traceID, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}


func (h *ContactHandler) updateContact(w http.ResponseWriter, r *http.Request) {
	
	traceID := r.Context().Value("traceID").(string)

	
	h.logger.Printf("[%s] Updating contact\n", traceID)

	
	var contact usecase.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		
		h.logger.Printf("[%s] Error decoding request body: %v\n", traceID, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	
	err = h.useCase.UpdateContact(r.Context(), &contact)
	if err != nil {
		
		h.logger.Printf("[%s] Error updating contact: %v\n", traceID, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}



func (h *ContactHandler) deleteContact(w http.ResponseWriter, r *http.Request) {
	
	traceID := r.Context().Value("traceID").(string)

	
	h.logger.Printf("[%s] Deleting contact\n", traceID)

	
	err := h.useCase.DeleteContact(r.Context())
	if err != nil {
		
		h.logger.Printf("[%s] Error deleting contact: %v\n", traceID, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}



type GroupHandler struct {
    useCase usecase.GroupUseCase
    logger  *log.Logger
}


func NewGroupHandler(useCase usecase.GroupUseCase, logger *log.Logger) *GroupHandler {
    return &GroupHandler{
        useCase: useCase,
        logger:  logger,
    }
}


func (h *GroupHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
    
    requestID := uuid.New()

    
    traceID := r.Header.Get("X-Trace-ID")
    if traceID == "" {
        traceID = uuid.New().String()
    }

    
    ctx := context.WithValue(r.Context(), "traceID", traceID)
    ctx = context.WithValue(ctx, "requestID", requestID)

    
    h.logger.Printf("[%s] %s %s %s\n", traceID, r.Method, r.URL.Path, r.Proto)

    
    switch r.Method {
    case http.MethodGet:
        h.logger.Println("GET /group")
        h.getGroup(w, r.WithContext(ctx))
    case http.MethodPost:
        h.logger.Println("POST /group")
        h.createGroup(w, r.WithContext(ctx))
    case http.MethodPut:
        h.logger.Println("PUT /group")
        h.updateGroup(w, r.WithContext(ctx))
    case http.MethodDelete:
        h.logger.Println("DELETE /group")
        h.deleteGroup(w, r.WithContext(ctx))
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}


func (h *GroupHandler) getGroup(w http.ResponseWriter, r *http.Request) {
    
    traceID := r.Context().Value("traceID").(string)

    
    h.logger.Printf("[%s] Getting group\n", traceID)

    
    requestID := r.Context().Value("requestID").(uuid.UUID)

    
    group, err := h.useCase.GetGroup(requestID)
    if err != nil {
        
        h.logger.Printf("[%s] Error getting group: %v\n", traceID, err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(group)
}


func (h *GroupHandler) createGroup(w http.ResponseWriter, r *http.Request) {
    
    traceID := r.Context().Value("traceID").(string)

    
    h.logger.Printf("[%s] Creating group\n", traceID)

    
    var group usecase.Group
    err := json.NewDecoder(r.Body).Decode(&group)
    if err != nil {
        
        h.logger.Printf("[%s] Error decoding request body: %v\n", traceID, err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    
    err = h.useCase.CreateGroup(&group)
    if err != nil {
        
        h.logger.Printf("[%s] Error creating group: %v\n", traceID, err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    
    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(group)
}


func (h *GroupHandler) updateGroup(w http.ResponseWriter, r *http.Request) {
    
    traceID := r.Context().Value("traceID").(string)

    
    h.logger.Printf("[%s] Updating group\n", traceID)

    
    var group usecase.Group
    err := json.NewDecoder(r.Body).Decode(&group)
    if err != nil {
        
        h.logger.Printf("[%s] Error decoding request body: %v\n", traceID, err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    
    err = h.useCase.UpdateGroup(&group)
    if err != nil {
        
        h.logger.Printf("[%s] Error updating group: %v\n", traceID, err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(group)
}


func (h *GroupHandler) deleteGroup(w http.ResponseWriter, r *http.Request) {
    
    traceID := r.Context().Value("traceID").(string)

    
    h.logger.Printf("[%s] Deleting group\n", traceID)

    
    err := h.useCase.DeleteGroup(r.Context())
    if err != nil {
        
        h.logger.Printf("[%s] Error deleting group: %v\n", traceID, err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    
    w.WriteHeader(http.StatusNoContent)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(nil)
}