require "constants"
require "tank"
require "shell"

lg = love.graphics

function love.load()
  SCREEN_W = 1000
  SCREEN_H = 800

  mouseX, mouseY = 0, 0

  VEL_DEBUG = false
  love.window.setMode(SCREEN_W, SCREEN_H)

  entities = {players={}, proj={}}  -- proj is projectiles

  Tank({num=1, type="tank"}, 200, 400)
end

function love.draw()
  for _, j in pairs(entities) do   -- Iterate each table in bodies.
    for x=1, #j do   -- For every body in that table, draw it.
      j[x]:draw()
    end
  end

  lg.setColor(1, 1, 1, 1)
  lg.print(love.timer.getFPS(), 10, 10)
end

function love.update(dt)
  mouseX, mouseY = love.mouse.getPosition()

  for _, j in pairs(entities) do
    for x=1, #j do
      j[x]:update(dt)
    end
  end
end

function love.keypressed(key)
  if key == "v" then
    VEL_DEBUG = not VEL_DEBUG
  end
end
