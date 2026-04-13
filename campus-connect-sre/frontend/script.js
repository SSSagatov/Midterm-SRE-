const API_BASE = "http://localhost:8080";
const postsContainer = document.getElementById("postsContainer");
const postForm = document.getElementById("postForm");
const formMessage = document.getElementById("formMessage");
const healthBadge = document.getElementById("healthBadge");
const refreshBtn = document.getElementById("refreshBtn");
const postCount = document.getElementById("postCount");

async function checkHealth() {
  try {
    const response = await fetch(`${API_BASE}/healthz`);
    if (!response.ok) throw new Error("Backend is unhealthy");
    const data = await response.json();
    healthBadge.textContent = data.status.toUpperCase();
    healthBadge.className = "ok";
  } catch (error) {
    healthBadge.textContent = "OFFLINE";
    healthBadge.className = "bad";
  }
}

async function loadPosts() {
  postsContainer.innerHTML = '<div class="empty">Loading posts...</div>';
  try {
    const response = await fetch(`${API_BASE}/api/posts`);
    if (!response.ok) throw new Error("Could not load posts");
    const posts = await response.json();
    renderPosts(posts);
  } catch (error) {
    postsContainer.innerHTML = `<div class="empty">Failed to load posts: ${error.message}</div>`;
  }
}

function renderPosts(posts) {
  postCount.textContent = `${posts.length} post${posts.length === 1 ? "" : "s"}`;
  if (!posts.length) {
    postsContainer.innerHTML = '<div class="empty">No posts yet. Create the first one.</div>';
    return;
  }

  postsContainer.innerHTML = posts
    .map(
      (post) => `
      <article class="post-card">
        <div class="post-meta">
          <span><strong>${escapeHtml(post.author)}</strong></span>
          <span>${new Date(post.created_at).toLocaleString()}</span>
        </div>
        <h3>${escapeHtml(post.title)}</h3>
        <p>${escapeHtml(post.content)}</p>
        <div class="like-row">
          <span>${post.likes} like${post.likes === 1 ? "" : "s"}</span>
          <button class="like-btn" data-id="${post.id}">Like post</button>
        </div>
      </article>
    `
    )
    .join("");

  document.querySelectorAll(".like-btn").forEach((button) => {
    button.addEventListener("click", async () => {
      const id = button.dataset.id;
      button.disabled = true;
      try {
        const response = await fetch(`${API_BASE}/api/posts/${id}/like`, { method: "POST" });
        if (!response.ok) throw new Error("Could not like post");
        await loadPosts();
      } catch (error) {
        alert(error.message);
      } finally {
        button.disabled = false;
      }
    });
  });
}

postForm.addEventListener("submit", async (event) => {
  event.preventDefault();
  formMessage.textContent = "Publishing...";
  formMessage.className = "message";

  const payload = {
    author: document.getElementById("author").value.trim(),
    title: document.getElementById("title").value.trim(),
    content: document.getElementById("content").value.trim(),
  };

  try {
    const response = await fetch(`${API_BASE}/api/posts`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });

    const data = await response.json();
    if (!response.ok) throw new Error(data.error || "Request failed");

    formMessage.textContent = "Post published successfully.";
    formMessage.className = "message ok";
    postForm.reset();
    await loadPosts();
  } catch (error) {
    formMessage.textContent = error.message;
    formMessage.className = "message bad";
  }
});

refreshBtn.addEventListener("click", loadPosts);

function escapeHtml(value) {
  return String(value)
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;")
    .replaceAll("'", "&#039;");
}

checkHealth();
loadPosts();
