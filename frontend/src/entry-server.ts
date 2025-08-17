// Fixed version of your render function with proper HTML structure handling
import { render as _render } from 'svelte/server'
import { compile } from 'mdsvex'
import App from './App.svelte'
import type { Artifact, Link } from './types'

function reviseImagePaths(artifact: Artifact, content: string) {
  if (!artifact.files) return content;
  artifact.files.forEach(file => {
    const relativePath = `./${file.name}`;
    const absolutePath = "/" + file.location;
    content = content.replace(
      new RegExp(`\\!\\[([^\\]]*)\\]\\(${relativePath.replace('.', '\\.')}\\)`, 'g'),
      `![$1](${absolutePath})`
    );
  });
  return content;
}

function processImageGrid(content: string) {
  const imageGridRegex = /__START_IMAGE_GRID__([\s\S]*?)__END_IMAGE_GRID__/g;
  content = content.replace(imageGridRegex, (match, content: string) => {
    const imageData = content.trim().split('\n')
      .filter(line => line.trim())
      .map(line => {
        const imageMatch = line.match(/!\[(.*?)\]\(([^)]+\.[a-zA-Z]+)\)/);
        return imageMatch ? { alt: imageMatch[1], src: imageMatch[2] } : null;
      })
      .filter(Boolean);
   
    const gridClass = imageData.length < 3 ? `image-grid col-${imageData.length}` : 'image-grid col-3';
   
    return `<div class="${gridClass}">
      ${imageData.map(img => `<img src="${img?.src}" alt="${img?.alt}" loading="lazy" />`).join('')}
    </div>`;
  });
  return content;
}

function processImage(content: string) {
  return content.replace(/!\[(.*?)\]\(([^)]+)\)/g, '<img src="$2" alt="$1" loading="lazy" />');
}

function processSideSections(content: string) {
  const sideParaRegex = /__START_PARA_WITH_SIDE_SECTION__([\s\S]*?)__END_PARA_WITH_SIDE_SECTION__/g;
  content = content.replace(sideParaRegex, (match, content: string) => {
    // extract side section content first
    const sectionHeadingMatch = content.match(/__START_SIDE_SECTION_HEADING__([\s\S]*?)__END_SIDE_SECTION_HEADING__/);
    const sectionListMatch = content.match(/__START_SIDE_SECTION_LIST__([\s\S]*?)__END_SIDE_SECTION_LIST__/);
    
    // remove side section markers from main content
    let mainContent = content
      .replace(/__START_SIDE_SECTION_HEADING__([\s\S]*?)__END_SIDE_SECTION_HEADING__/, '')
      .replace(/__START_SIDE_SECTION_LIST__([\s\S]*?)__END_SIDE_SECTION_LIST__/, '')
      .trim();

    // process any regular sections within the main content
    mainContent = mainContent.replace(
      /__START_SECTION_HEADING__([\s\S]*?)__END_SECTION_HEADING__[\s\S]*?__START_SECTION_LIST__([\s\S]*?)__END_SECTION_LIST__/g,
      (sectionMatch, heading, list) => {
        // ensure proper list formatting by adding blank lines around list items
        const formattedList = '\n\n' + list.trim() + '\n\n';
        return `\n\n<!-- SECTION_CARD_START -->
<!-- SECTION_HEADING_START -->${heading.trim()}<!-- SECTION_HEADING_END -->
<!-- SECTION_LIST_START -->${formattedList}<!-- SECTION_LIST_END -->
<!-- SECTION_CARD_END -->\n\n`;
      }
    );
    
    const sectionHeading = sectionHeadingMatch ? sectionHeadingMatch[1].trim() : '';
    let sectionList = sectionListMatch ? sectionListMatch[1].trim() : '';
    
    // Ensure proper list formatting by adding blank lines around list items
    if (sectionList) {
      sectionList = '\n\n' + sectionList + '\n\n';
    }
    
    return `\n\n<!-- FLEX_CONTAINER_START -->
<!-- MAIN_CONTENT_START -->
${mainContent}
<!-- MAIN_CONTENT_END -->
<!-- SIDE_HEADING_START -->${sectionHeading}<!-- SIDE_HEADING_END -->
<!-- SIDE_LIST_START -->${sectionList}<!-- SIDE_LIST_END -->
<!-- FLEX_CONTAINER_END -->\n\n`;
  });

  return content;
}

function processSections(content: string) {
  const standaloneSectionRegex = /__START_SECTION_HEADING__([\s\S]*?)__END_SECTION_HEADING__[\s\S]*?__START_SECTION_LIST__([\s\S]*?)__END_SECTION_LIST__/g;
  content = content.replace(standaloneSectionRegex, (match, heading, list) => {
    // ensure proper list formatting by adding blank lines around list items
    const formattedList = '\n\n' + list.trim() + '\n\n';
    return `\n\n<!-- SECTION_CARD_START -->
<!-- SECTION_HEADING_START -->${heading.trim()}<!-- SECTION_HEADING_END -->
<!-- SECTION_LIST_START -->${formattedList}<!-- SECTION_LIST_END -->
<!-- SECTION_CARD_END -->\n\n`;
  });
  return content;
}

function addIdsToHeadings(content: string) {
  const headingRegex = /<h([1-6])>(.*?)<\/h[1-6]>/g;
  content = content.replace(headingRegex, (match, level, text) => {
    const id = text.toLowerCase()
      .replace(/[^a-z0-9\s]/g, '')
      .replace(/\s+/g, '-')
      .trim();
    return `<h${level} id="${id}"><a href="#${id}">${text}</a></h${level}>`;
  });
  return content;
}

function openLinksInNewTabs(content: string) {
  const linkRegex = /<a\s+([^>]*?)href=["']([^"']*?)["']([^>]*?)>/gi;
  content = content.replace(linkRegex, (match, beforeHref, href, afterHref) => {
    const isExternal = href.startsWith('http') || href.startsWith('//');
    const hasTarget = /target\s*=/i.test(beforeHref + afterHref);
    
    if (isExternal && !hasTarget) {
      return `<a ${beforeHref}href="${href}" target="_blank"${afterHref}>`;
    }
    return match;
  });
  return content;
}

function styleSectionCard(content: string) {
  content = content.replace(
    /<!-- SECTION_CARD_START -->([\s\S]*?)<!-- SECTION_CARD_END -->/g,
    (match, content) => {
      let sectionHeading = content.match(/<!-- SECTION_HEADING_START -->([\s\S]*?)<!-- SECTION_HEADING_END -->/)?.[1] || '';
      const sectionList = content.match(/<!-- SECTION_LIST_START -->([\s\S]*?)<!-- SECTION_LIST_END -->/)?.[1] || '';
      
      // clean up heading HTML by removing unnecessary <p> tags
      sectionHeading = sectionHeading.replace(/<\/?p>/g, '').trim();
      
      return `<section class="border-2 border-vaxitas-tertiary p-4 rounded-lg relative my-4 card-section">
        <div class="absolute top-0 -translate-y-1/2 my-0 bg-vaxitas-primary text-vaxitas-secondary font-semibold px-2 left-2 text-lg">${sectionHeading}</div>
        ${sectionList.trim()}
      </section>`;
    }
  );
  return content;
}

function styleSideSectionFlexContainer(content: string) {
  content = content.replace(
    /<!-- FLEX_CONTAINER_START -->([\s\S]*?)<!-- FLEX_CONTAINER_END -->/g,
    (match, content) => {
      const mainContent = content.match(/<!-- MAIN_CONTENT_START -->([\s\S]*?)<!-- MAIN_CONTENT_END -->/)?.[1] || '';
      let sideHeading = content.match(/<!-- SIDE_HEADING_START -->([\s\S]*?)<!-- SIDE_HEADING_END -->/)?.[1] || '';
      const sideList = content.match(/<!-- SIDE_LIST_START -->([\s\S]*?)<!-- SIDE_LIST_END -->/)?.[1] || '';

      // clean up heading HTML by removing unnecessary <p> tags
      sideHeading = sideHeading.replace(/<\/?p>/g, '').trim();
      
      return `<div class="flex flex-col lg:flex-row gap-2 lg:gap-8 items-start">
        <div class="flex-1">
          ${mainContent.trim()}
        </div>
        <section class="w-full lg:w-1/2 lg:max-w-lg border-2 border-vaxitas-tertiary p-4 rounded-lg relative mx-auto lg:mx-0 card-section mb-4">
          <div class="absolute top-0 -translate-y-1/2 my-0 bg-vaxitas-primary text-vaxitas-secondary font-semibold px-2 left-2 text-lg">${sideHeading}</div>
          ${sideList.trim()}
        </section>
      </div>`;
    }
  );
  return content;
}

function extractFootnotes(content: string) {
  const footnotes: { [key: string]: string } = {};
  let processedContent = content;
  
  const lines = content.split('\n');
  const resultLines: string[] = [];
  
  let currentFootnoteKey = '';
  let currentFootnoteContent: string[] = [];
  let inFootnote = false;
  
  for (let i = 0; i < lines.length; i++) {
    const line = lines[i];
    
    // check if this line starts a new footnote
    const footnoteStart = line.match(/^\[(\^[^\]]+)\]:\s*(.*)/);
    
    if (footnoteStart) {
      // save previous footnote if exists
      if (currentFootnoteKey && currentFootnoteContent.length > 0) {
        footnotes[currentFootnoteKey] = currentFootnoteContent.join('\n').trim();
      }
      
      // start new footnote
      currentFootnoteKey = footnoteStart[1];
      currentFootnoteContent = [footnoteStart[2]]; // start with content after the colon
      inFootnote = true;
      
      continue;
    }
    
    // if we're in a footnote, check if this line continues it
    if (inFootnote) {
      // check if this line starts a new heading
      const isNewHeading = line.match(/^#{1,6}\s/);
      
      if (isNewHeading) {
        // handle the end of the current footnote by adding its content to the mapping
        if (currentFootnoteKey && currentFootnoteContent.length > 0) {
          footnotes[currentFootnoteKey] = currentFootnoteContent.join('\n').trim();
        }
        
        inFootnote = false;
        currentFootnoteKey = '';
        currentFootnoteContent = [];
        
        resultLines.push(line);
        
        continue;
      }
      
      currentFootnoteContent.push(line);
      continue;
    }
    
    resultLines.push(line);
  }
  
  // save the last footnote if we were processing one
  if (inFootnote && currentFootnoteKey && currentFootnoteContent.length > 0) {
    footnotes[currentFootnoteKey] = currentFootnoteContent.join('\n').trim();
  }
  
  processedContent = resultLines.join('\n').trim();
  
  return {
    content: processedContent,
    footnotes
  };
}

function processInlineReferences(content: string, footnotes: { [key: string]: string }) {
  return content.replace(/\[(\^[^\]]+)\]/g, (match, key) => {
    if (footnotes[key]) {
      const id = key.replace('^', ''); // remove the ^ to get just the footnote identifier
      return `<sup><a href="#footnote-${id}" id="ref-${id}" class="text-vaxitas-secondary hover:underline">${id}</a></sup>`;
    }
    return match;
  });
}

function processFootnoteParagraphs(content: string) {
  const htmlContent = content
    .split('\n\n')
    .map(para => para.trim())
    .filter(para => para.length > 0)
    .map((para, idx, arr) => {
      // we don't want to wrap in <p> if it's an image or image grid
      if (para.match(/^<img|^<div class="image-grid/)) return para;
      
      // apply basic markdown formatting for text paragraphs
      const formatted = para
        .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
        .replace(/\*(.*?)\*/g, '<em>$1</em>')
        .replace(/`(.*?)`/g, '<code>$1</code>');
      
      return `<p class=${idx === arr.length - 1 ? "inline" : ""}>${formatted}</p>`;
    })
    .join('');
  return htmlContent;
}

function wrapFootnoteInListItem(content: string, id: string) {
  return `<li id="footnote-${id}" class="text-sm">${content} <a href="#ref-${id}" class="text-vaxitas-secondary hover:underline">↩</a></li>`;
}

function appendFootnotes(htmlContent: string, footnotes: { [key: string]: string }, artifact: Artifact) {
  const hasFootnotes = Object.keys(footnotes).length > 0;
  if (!hasFootnotes) return htmlContent;
  
  let footnoteHtml = '<ol class="space-y-2 mt-4">';
  
  for (const [key, content] of Object.entries(footnotes)) {
    const id = key.replace('^', '');
    // process the footnote content similar to main content processing
    let processedContent = content;
    processedContent = reviseImagePaths(artifact, processedContent);
    processedContent = processImageGrid(processedContent);
    processedContent = processImage(processedContent);
    const htmlContent = processFootnoteParagraphs(processedContent);
    footnoteHtml += wrapFootnoteInListItem(htmlContent, id);
  }
  
  footnoteHtml += '</ol>';
  
  return htmlContent + footnoteHtml;
}

function processNowPageContent(url: string, content: string, secondaryNavData: Link) {
  if (url !== "/now") return content;
  // check if we have sat least one item in secondary nav data
  if (!(secondaryNavData.items && secondaryNavData.items.length > 0)) return content;
  
  const targetText = "Click here for the most recent monthly log.";
  const firstItem = secondaryNavData.items[0];
  const linkHtml = `Click <a href="${firstItem.path}">here</a> for the most recent monthly log.`;
  return content.replace(targetText, linkHtml);
}

export async function render(
  url: string, 
  primaryNavData: Link[],
  secondaryNavData: Link, 
  artifact: Artifact,
) {
  let compiledArtifact = artifact;

  if (!artifact.content) return;

  try {
    let processedContent = artifact.content;

    // step 1: extract footnotes BEFORE any other processing
    const footnoteResult = extractFootnotes(processedContent);
    processedContent = footnoteResult.content;
    const footnotes = footnoteResult.footnotes;

    // step 2: process inline footnote references IMMEDIATELY after extraction
    processedContent = processInlineReferences(processedContent, footnotes);

    // step 3: add link to the most recent monthly log in the index page
    processedContent = processNowPageContent(url, processedContent, secondaryNavData);

    // step 4: handle image paths
    processedContent = reviseImagePaths(artifact, processedContent);

    // step 5: handle image grids
    processedContent = processImageGrid(processedContent);

    // step 6: process side sections
    processedContent = processSideSections(processedContent);

    // step 7: process regular sections
    processedContent = processSections(processedContent);

    // step 8: compile with mdsvex
    const result = await compile(processedContent, {
      extensions: ['.svx', '.md'],
      smartypants: true,
    });
    let compiledContent = result?.code || processedContent;

    // step 9: add IDs to headings
    compiledContent = addIdsToHeadings(compiledContent);

    // step 10: add target="_blank" to external links
    compiledContent = openLinksInNewTabs(compiledContent);

    // step 11: format the section cards
    compiledContent = styleSectionCard(compiledContent);

    // step 12: format the containers with side sections
    compiledContent = styleSideSectionFlexContainer(compiledContent);

    // compiledContent = compiledContent.replace(
    //   /{@html `([^`]*)`}/g, 
    //   (match, htmlContent) => {
    //     // Unescape the HTML content
    //     return htmlContent
    //       .replace(/\\`/g, '`')
    //       .replace(/\\\$/g, '$')
    //       .replace(/\\\\/g, '\\');
    //   }
    // );

    // step 13: add footnotes section
    compiledContent = appendFootnotes(compiledContent, footnotes, artifact);

    compiledArtifact = {
      ...artifact,
      content: compiledContent
    };
  } catch (error) {
    console.error('Error compiling markdown:', error);
    compiledArtifact = {
      ...artifact,
      content: artifact.content
    };
  }

  return [
    _render(App, {
      props: {
        url,
        primaryNavData,
        secondaryNavData,
        artifact: compiledArtifact,
      }
    }), 
    compiledArtifact
  ];
}
